package controllers

import (
	"golangRest/src/database"
	"golangRest/src/models"
	"golangRest/src/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewStockController struct {
	repo *repository.StockRepository
}

func NewStockControllerInstance() *NewStockController {
	repo := repository.NewStockRepository(database.DB)
	return &NewStockController{repo: repo}
}

func (sr *NewStockController) GetAllData(c *gin.Context) {
	var responsedata models.StockModel
	if err := sr.repo.DB.Table("stock").Find(&responsedata).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mesdage": "json a",
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"mesdage": "json a",
		"data":    responsedata,
	})

}

func (sr *NewStockController) Insert(c *gin.Context) {

}

func (sr *NewStockController) ShowData(c *gin.Context) {

}

func (sr *NewStockController) Delete(c *gin.Context) {

}
func (sr *NewStockController) Show(c *gin.Context) {

}
