package controllers

import (
	"golangRest/src/database"
	"golangRest/src/repository"

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

}

func (sr *NewStockController) Insert(c *gin.Context) {

}

func (sr *NewStockController) ShowData(c *gin.Context) {

}

func (sr *NewStockController) Delete(c *gin.Context) {

}
func (sr *NewStockController) Show(c *gin.Context) {

}
