package handler

import (
	"booking/config"
	"booking/models"
	"booking/pkg/helper/helper"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


// Register godoc
// @Summary Foydalanuvchini ro‘yxatdan o‘tkazish
// @Description Yangi foydalanuvchini yaratadi
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body models.CreateUser true "Ro‘yxatdan o‘tish ma'lumotlari"
// @Success 201 {object} models.User
// @Failure 400 {object} models.DefaultError
// @Failure 409 {object} models.DefaultError
// @Failure 500 {object} models.DefaultError
// @Router /register [post]
func (h *Handler) Register(c *gin.Context) {
	var createUser models.CreateUser

	err := c.ShouldBindJSON(&createUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Error parsing registration data: " + err.Error(),
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Error hashing password: " + err.Error(),
		})
		return
	}

	createUser.Password = string(hashedPassword)

	role, err := h.strg.Role().GetByName(context.Background(), &models.Role{Name: "customer"})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Role not found: " + err.Error(),
		})
		return
	}

	createUser.RoleId = &role.Id

	userId, err := h.strg.User().Create(context.Background(), &createUser)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			c.JSON(http.StatusConflict, models.DefaultError{
				Message: "User already exists, please login!",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Error creating user: " + err.Error(),
		})
		return
	}

	user, err := h.strg.User().GetByID(context.Background(), userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Error fetching user information: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}


// Login godoc
// @Summary Foydalanuvchini tizimga kiritish
// @Description Foydalanuvchi login va parol orqali tizimga kiradi
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body models.Login true "Login ma'lumotlari"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} models.DefaultError
// @Failure 401 {object} models.DefaultError
// @Failure 500 {object} models.DefaultError
// @Router /login [post]
func (h *Handler) Login(c *gin.Context) {
	var login models.Login

	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{Message: "Error parsing login data: " + err.Error()})
		return
	}

	resp, err := h.strg.User().GetByUserName(context.Background(), &login)
	if err != nil {
		if err.Error() == "no rows in result set" {
			c.JSON(http.StatusBadRequest, models.DefaultError{Message: "User not found, please register first"})
			return
		}
		c.JSON(http.StatusInternalServerError, models.DefaultError{Message: "Error fetching user data: " + err.Error()})
		return
	}

	if resp.Password == "" {
		c.JSON(http.StatusUnauthorized, models.DefaultError{Message: "Invalid credentials"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(resp.Password), []byte(login.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.DefaultError{Message: "Invalid credentials"})
		return
	}

	role, err := h.strg.Role().GetByID(context.Background(), &models.PrimaryKey{Id: *resp.RoleId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{Message: "Error fetching role data: " + err.Error()})
		return
	}
	fmt.Println(role)

	data := map[string]interface{}{
		"id":         resp.Id,
		"first_name": resp.FirstName,
		"last_name":  resp.LastName,
		"username":   resp.UserName,
		"email":      resp.Email,
		"role_id":    resp.RoleId,
		"created_at": resp.CreatedAt,
		"updated_at": resp.UpdatedAt,
		"role":       role.Name,
	}
	token, err := helper.GenerateJWT(data, config.TimeExpiredAt, h.cfg.SekretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{Message: "Error generating JWT token: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.LoginResponse{Token: token, UserData: resp})
}
