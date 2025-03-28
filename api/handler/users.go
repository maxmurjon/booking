package handler

import (
	"booking/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
