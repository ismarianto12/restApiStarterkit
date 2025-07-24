package controllers

import (
	"fmt"
	"golangRest/src/database"
	"golangRest/src/models"
	"golangRest/src/repository"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
func (ptk *BarangController) Delet(c *gin.Context) {
	id := c.Param("id")
	idF, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error page", "version": 12})
		return
	}
	// fmt.Println("Pantek %s", idF)
	err = ptk.repo.DB.Exec("delete from barang where id = ? ", idF).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error page", "version": 12})
		return
	}
	c.JSON(http.StatusOK, gin.H{"messages": "berhasil menghapus data barang"})
}

func (pt *BarangController) UpdateData(c *gin.Context) {

	var barang models.BarangModel

	if err := c.ShouldBindJSON(&barang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "gagal payload",
			"data":    nil,
		})
		return
	}

	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id tidak valid",
		})
		return
	}

	var existingBarang models.BarangModel
	if err := pt.repo.DB.First(&existingBarang, idInt).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "barang tidak ditemukan",
		})
		return
	}

	existingBarang.Kode = barang.Kode
	existingBarang.Nama = barang.Nama
	existingBarang.KategoryId = barang.KategoryId
	existingBarang.Deskripsi = barang.Deskripsi
	existingBarang.Gambar = barang.Gambar

	if err := pt.repo.DB.Find(&barang, idInt).Save(existingBarang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "gagal update",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil update",
		"data":    existingBarang,
	})

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

func (ptk *BarangController) GetSemuaKontol(c *gin.Context) {

	var barang []models.BarangModel
	err := godotenv.Load()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error page", "version": 12})
		return
	}
	appversion := os.Getenv("APP_VERSION")
	if c.Query("page") == "" || c.Query("total") == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error page", "version": appversion})
		return
	}
	if err := ptk.repo.DB.Raw("select * from barang").Scan(&barang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "bangsaat error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":     barang,
		"status":   "ok",
		"messages": "success ok",
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

func (ptk *BarangController) GetBarangByid(c *gin.Context) {
	var barang models.BarangModel
	id := c.Param("id")
	idIht, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID tidak valid",
		})
		return
	}
	if err := ptk.repo.DB.First(&barang, idIht).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    barang,
			"message": "successfull data",
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"data":    barang,
		"message": "successfull data",
	})

}

// chec untuk PUT /barang/update/:id
// func (bc *BarangController) UpdateData(c *gin.Context) {
// 	id := c.Param("id")
// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "Data dengan ID " + id + " berhasil diupdate (simulasi)",
// 	})
// }
