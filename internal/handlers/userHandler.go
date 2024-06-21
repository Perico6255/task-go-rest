package handlers

import (
	"Perico6255/task-go-rest/internal/models"
	"Perico6255/task-go-rest/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	s services.UserService
}

func NewUserHandler(s services.UserService) *UserHandler {
	return &UserHandler{s}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.UserModel
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	createdUser, err := h.s.Create(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, createdUser)
}
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.s.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, users)
	return
}
