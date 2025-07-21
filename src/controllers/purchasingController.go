package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PurchaseRequest struct {
	ProductID int    `json:"product_id"`
	Qty       int    `json:"qty"`
	Note      string `json:"note"`
}

type PurchaseResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type NewPurchasingController struct {
}

func NewPurchasingControllerInstance() *NewPurchasingController {
	return &NewPurchasingController{}
}

func (uc *NewPurchasingController) GetAllData(c *gin.Context) {
	var req PurchaseRequest
	c.ShouldBindJSON(&req)
	response := PurchaseResponse{
		Status:  "succes",
		Message: req.Note,
	}
	c.JSON(http.StatusOK, response)
}

func (pantek *NewPurchasingController) Store(c *gin.Context) {
	// c.ShouldBindJSON(&c.req)

	c.JSON(http.StatusOK, gin.H{
		"testing": "test",
	})

}
