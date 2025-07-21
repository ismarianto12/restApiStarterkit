package controllers

import (
	"golangRest/src/database"
	"golangRest/src/models"
	"golangRest/src/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	repo *repository.UserRepository
}

type UserResponse struct {
	Data    []models.UserMode `json:"data"`
	Status  string            `json:"status"`
	Message string            `json:"message"`
}

func NewUserController() *UserController {
	repo := repository.NewUserRepository(database.DB) // ✅ kirim database.DB sebagai argumen
	return &UserController{repo: repo}
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.repo.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	custom := UserResponse{
		Data:    users,
		Status:  "200",
		Message: "data berhasil",
	}
	c.JSON(http.StatusOK, custom)
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.UserMode

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := uc.repo.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (uc *UserController) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := uc.repo.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
