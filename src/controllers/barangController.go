package controllers

import (
	"fmt"
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

type BarangController struct{} // Tidak perlu repository

func NewBarangController() *BarangController {
	return &BarangController{}
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
