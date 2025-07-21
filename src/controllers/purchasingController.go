package controllers

import (
	"golangRest/src/database"
	"golangRest/src/models"
	"golangRest/src/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PurchaseRequest struct {
	ProductID int    `json:"product_id"`
	Qty       int    `json:"qty"`
	Note      string `json:"note"`
}

type PurchaseResponse struct {
	Data    []models.PurcahsingModel `json:"data"`
	Status  string                   `json:"status"`
	Message string                   `json:"message"`
}

type NewPurchasingController struct {
	repo *repository.PurcashingRepository
}

func NewPurchasingControllerInstance() *NewPurchasingController {
	repo := repository.NewPurcashingRepository(database.DB)
	return &NewPurchasingController{repo: repo}
}

func (uc *NewPurchasingController) GetAllData(c *gin.Context) {
	var req PurchaseRequest
	prdata, err := uc.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve purhcasing data"})
		return
	}
	c.ShouldBindJSON(&req)
	response := PurchaseResponse{
		Data:    []models.PurcahsingModel{*prdata},
		Status:  "succes",
		Message: req.Note,
	}
	c.JSON(http.StatusOK, response)
}

func (pantek *NewPurchasingController) Store(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"testing": "test",
	})
}

func (ctrl *NewPurchasingController) getAlldat(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"testing": "test",
	})
}
