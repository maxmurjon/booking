package handler

import (
	"booking/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new user
// @Description This endpoint creates a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param user body models.CreateUser true "User data"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.DefaultError
// @Failure 500 {object} models.DefaultError
// @Router /createuser [post]
func (h *Handler) CreateUser(c *gin.Context) {
	var entity *models.CreateUser
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	id, err := h.strg.User().Create(context.Background(), entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to create user: " + err.Error(),
		})
		return
	}

	user, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Id: id.Id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve created user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "User has been created",
		Data:    user,
	})
}

// @Summary Update an existing user
// @Description This endpoint updates user information
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param user body models.UpdateUser true "User data"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.DefaultError
// @Failure 500 {object} models.DefaultError
// @Router /updateuser [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	var entity models.UpdateUser
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Yaroqsiz so'rov tanasi: " + err.Error(),
		})
		return
	}

	if entity.Id == "" {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Foydalanuvchi ID-si talab qilinadi",
		})
		return
	}

	if _, err := h.strg.User().Update(context.Background(), &entity); err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Foydalanuvchini yangilashda xato: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Foydalanuvchi yangilandi",
		Data:    entity.Id,
	})
}

// @Summary Get a list of users
// @Description This endpoint retrieves a list of users
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Success 200 {object} models.GetListUserResponse
// @Failure 500 {object} models.DefaultError
// @Router /users [get]
func (h *Handler) GetUsersList(c *gin.Context) {
	resp, err := h.strg.User().GetList(context.Background(), &models.GetListUserRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve user list: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary Get a user by ID
// @Description This endpoint retrieves a user by their ID
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param id path string true "User ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 404 {object} models.DefaultError
// @Router /user/{id} [get]
func (h *Handler) GetUsersByIDHandler(c *gin.Context) {
	id := c.Param("id")

	user, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Id: &id})
	if err != nil {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "User not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OK",
		Data:    user,
	})
}


// @Summary Delete a user
// @Description This endpoint deletes a user by their ID
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param id path string true "User ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.DefaultError
// @Router /deleteuser/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	deletedUser, err := h.strg.User().Delete(context.Background(), &models.UserPrimaryKey{Id: &id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to delete user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "User has been deleted",
		Data:    deletedUser,
	})
}
