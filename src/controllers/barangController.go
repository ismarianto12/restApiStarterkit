package controllers

import (
	"fmt"
	"golangRest/src/database"
	"golangRest/src/models"
	"golangRest/src/repository"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type Barang struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

type BarangController struct {
	repo *repository.BarangRepository
}

func NewBarangController() *BarangController {
	repo := repository.NewBarangRepository(database.DB)
	return &BarangController{repo: repo}
}

func (bc *BarangController) UploadDfile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no file"})
	}

	allowedExts := []string{".jpg", ".jpeg", ".png", ".pdf"}
	fileExt := strings.ToLower(filepath.Ext(file.Filename))
	isValid := false
	for _, ext := range allowedExts {
		if fileExt == ext {
			isValid = true
			break
		}
	}
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("File extension '%s' is not allowed", fileExt),
		})
		return
	}
	// file.ex
	dst := filepath.Join("uploads", filepath.Base(file.Filename))
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save the file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully!", "filename": file.Filename})

}

func (bc *BarangController) GetAllData(c *gin.Context) {
	barangList := []Barang{
		{ID: 1, Name: "Laptop", Price: 10000000, Stock: 10},
		{ID: 2, Name: "Mouse", Price: 150000, Stock: 50},
		{ID: 3, Name: "Keyboard", Price: 300000, Stock: 30},
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    barangList,
	})
}

func (ptk *BarangController) GetSemuaKontol(c *gin.Context, *models.BarangModel) {
	var barang models.BarangModel
	if  barang,err := c.ShouldBindJSON(barang) == nil {

	}


	barang, err := ptk.repo.GetAllBarang(2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "bangsaat error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": barang, "message": "success",
	})

}

func (ptk *BarangController) Store(c *gin.Context) {
	var barang models.BarangModel
	// c.BindJSON(barang)
	if err := c.ShouldBindJSON(&barang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "istnot json data"})
		return
	}
	ptk.repo.InsertData(&barang)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    barang,
	})

}

// Handler untuk POST /barang/create
func (bc *BarangController) CreateData(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Data berhasil dibuat (simulasi)",
	})
}

// Handler untuk PUT /barang/update/:id
func (bc *BarangController) UpdateData(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data dengan ID " + id + " berhasil diupdate (simulasi)",
	})
}
