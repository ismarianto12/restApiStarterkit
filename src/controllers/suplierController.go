package controllers

import (
	"golangRest/src/database"
	"golangRest/src/models"
	"golangRest/src/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuplierController struct {
	repo *repository.SuplierRepository
}

func NewSuplierControllerInstance() *SuplierController {
	repopath := repository.NewSuplierRepositoryInstance(database.DB)
	return &SuplierController{repo: repopath}

}

func (ptk *SuplierController) Index(c *gin.Context) {
	response, err := ptk.repo.GetallData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messages": "not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "success load data",
		"data":    response,
	})

}
func (ptk *SuplierController) CreateData(c *gin.Context) {
	var supplier models.SuplierModel

	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"status":  400,
			"message": err.Error(),
		})
		return
	}

	err := ptk.repo.DB.Create(supplier).Error
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"status":  400,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "success load data",
	})

}
