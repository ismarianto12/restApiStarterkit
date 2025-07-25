package controllers

import (
	"errors"
	"fmt"
	"golangRest/src/database"
	"golangRest/src/models"
	"golangRest/src/repository"
	"golangRest/src/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	repo := repository.NewUserRepository(database.DB)
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

func (ptj *UserController) ShowUser(c *gin.Context) {
	var user models.UserMode
	id := c.Param("id")
	fid, err := strconv.Atoi(id)
	fmt.Print(fid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "id": fid})
		return
	}

	erp := ptj.repo.DB.Where("id", fid).First(&user)
	if erp != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "data kosong"})
			return
		}

	}
	c.JSON(http.StatusOK, gin.H{"data": user, "message": "data berhasil di temukan"})

}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.UserMode

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hasHedPassword, err := utils.Haspassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ecrypt data"})
		return
	}
	user.Password = hasHedPassword

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

func (auth *UserController) Login(c *gin.Context) {
	var user models.UserMode
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request tidak valid"})
		return
	}
	// passnya, errc := utils.Haspassword(user.Password)
	// if errc != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid has password tidak valid"})
	// 	return
	// }
	resp, err := auth.repo.LoginAuth(user.Username, user.Password)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan saat login"})
		}
		return
	}
	token, errtoket := utils.GenerateToken(resp.Username, resp.LevelId)
	if errtoket != nil {
		fmt.Printf("%+v\n", errtoket)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token", "err": errtoket})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": 200, "username": resp.Username, "token": token, "messages": "login berhasil authentikasi ."})

}

func EncpDe(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization header required",
			"code":  http.StatusUnauthorized,
		})

	}
	// bearer prefix
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	claims, err := utils.DecodeToken(tokenString)
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusOK, gin.H{
			"data": "crot daalam", "erro": err.Error(),
		})
	}
	// if userID, ok := claims.LevelId; ok {
	// 	fmt.Println("User ID:", userID)
	// }
	c.JSON(http.StatusOK, gin.H{
		"data":    claims.LevelID,
		"message": "success loaded data claim paramaeter",
	})

}
