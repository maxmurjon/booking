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

func (h *Handler) Register(c *gin.Context) {
	var createUser models.CreateUser

	// JSONni bind qilish
	err := c.ShouldBindJSON(&createUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Error parsing registration data: " + err.Error(),
		})
		return
	}

	// Parolni hash qilish
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Error hashing password: " + err.Error(),
		})
		return
	}

	createUser.Password = string(hashedPassword)

	// Role ID ni olish
	role, err := h.strg.Role().GetByName(context.Background(), &models.Role{Name: "customer"})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Role not found: " + err.Error(),
		})
		return
	}
	fmt.Println(role)
	createUser.RoleId = &role.Id
	fmt.Println(createUser, role)
	// Foydalanuvchini yaratish
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

	// Foydalanuvchi ma'lumotlarini olish
	user, err := h.strg.User().GetByID(context.Background(), userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Error fetching user information: " + err.Error(),
		})
		return
	}

	// Yaratilgan foydalanuvchini qaytarish
	c.JSON(http.StatusCreated, user)
}

func (h *Handler) Login(c *gin.Context) {
	var login models.Login

	// JSONni bind qilish
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{Message: "Error parsing login data: " + err.Error()})
		return
	}

	// Foydalanuvchini username bo‘yicha olish
	resp, err := h.strg.User().GetByUserName(context.Background(), &login)
	if err != nil {
		if err.Error() == "no rows in result set" {
			c.JSON(http.StatusBadRequest, models.DefaultError{Message: "User not found, please register first"})
			return
		}
		c.JSON(http.StatusInternalServerError, models.DefaultError{Message: "Error fetching user data: " + err.Error()})
		return
	}
	// Parolni tekshirish
	if resp.Password == "" {
		c.JSON(http.StatusUnauthorized, models.DefaultError{Message: "Invalid credentials"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(resp.Password), []byte(login.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.DefaultError{Message: "Invalid credentials"})
		return
	}

	// Role ID ni olish
	role, err := h.strg.Role().GetByID(context.Background(), &models.PrimaryKey{Id: *resp.RoleId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{Message: "Error fetching role data: " + err.Error()})
		return
	}
	fmt.Println(role)
	// JWT token yaratish
	data := map[string]interface{}{
		"id":         resp.Id,
		"first_name": resp.FirstName,
		"last_name":  resp.LastName,
		"username":   resp.UserName,
		"email":      resp.Email,
		"role_id":    resp.RoleId,
		"created_at": resp.CreatedAt,
		"updated_at": resp.UpdatedAt,
		"role":       role.Id,
	}
	token, err := helper.GenerateJWT(data, config.TimeExpiredAt, h.cfg.SekretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{Message: "Error generating JWT token: " + err.Error()})
		return
	}

	// JWT token va foydalanuvchi ma'lumotlarini qaytarish
	c.JSON(http.StatusOK, models.LoginResponse{Token: token, UserData: resp})
}
