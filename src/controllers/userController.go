package controllers

import (
	"golangRest/src/models"
	"golangRest/src/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	repo *repository.UserRepository
}

// NewUserController creates a new UserController instance
func NewUserController() *UserController {
	return &UserController{
		repo: repository.NewUserRepository(),
	}
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	// Retrieve all users from the repository
	users, err := uc.repo.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUser handles POST requests for creating a new user
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
