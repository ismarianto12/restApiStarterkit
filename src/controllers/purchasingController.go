package controllers

import (
	"fmt"
	"golangRest/src/database"
	"golangRest/src/models"
	"golangRest/src/repository"
	"net/http"
	"strconv"

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
	c.JSON(http.StatusOK, gin.H{
		"data": prdata,
	})
}

func (pantek *NewPurchasingController) Store(c *gin.Context) {
	var purcahsing models.PurcahsingModel
	err := pantek.repo.Create(&purcahsing)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "ngewe pakai cinta",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"testing": "test",
	})
}

func (sh *NewPurchasingController) Show(c *gin.Context) {
	id := c.Param("id")
	idPamaram, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}
	//get data store fromdb
	var purcahsing models.PurcahsingModel
	if err := sh.repo.DB.Find(&purcahsing, idPamaram).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data":     err,
			"messages": "pantek kehidupan kontol",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":     purcahsing,
		"messages": "data berhasil di load",
	})

}

func (edit *NewPurchasingController) Edit(c *gin.Context) {

	var purcahsing models.PurcahsingModel
	fmt.Println("barang " + purcahsing.Pemeriksa)

	id := c.Param("id")
	clId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"testing": "test",
		})
	}
	var existingPurcahsing models.PurcahsingModel
	if err := edit.repo.DB.First(&existingPurcahsing, clId).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": err,
		})
	}
	//existing data
	existingPurcahsing.FactureCode = purcahsing.FactureCode
	existingPurcahsing.Pemeriksa = purcahsing.Pemeriksa

	if err := edit.repo.DB.Find(&purcahsing, clId).Save(existingPurcahsing).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": "error can\t update data ",
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"data": purcahsing,
	})
}

func (ctrl *NewPurchasingController) GetAlldat(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"testing": "test",
	})
}
