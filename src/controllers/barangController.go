package controllers

import (
	"net/http"

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

// Handler untuk GET /barang/list
func (bc *BarangController) GetAllData(c *gin.Context) {
	// Data statis (hardcoded)
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
