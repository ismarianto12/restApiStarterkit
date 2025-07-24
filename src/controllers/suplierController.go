package controllers

import (
	"golangRest/src/database"
	"golangRest/src/models"
	"golangRest/src/repository"
	"net/http"
	"strconv"

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

func (ptk *SuplierController) UpdateData(c *gin.Context) {
	var existingbarang models.SuplierModel
	var barangpayload models.SuplierModel

	id := c.Param("id")
	Fid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "not valid id",
		})
		return
	}

	if efrr := c.ShouldBindJSON(&existingbarang).Error; efrr != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "not valid id",
		})
		return
	}

	//find data
	if errd := ptk.repo.DB.First(&id, Fid).Scan(&existingbarang).Error; errd != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "not valid id",
		})
		return
	}

	existingbarang.Address = barangpayload.Address
	existingbarang.Nama = barangpayload.Nama
	existingbarang.Contact = barangpayload.Contact

	if errd := ptk.repo.DB.Save(existingbarang).Error; errd != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "not valid id",
		})
		return
	}

}

func (ptk *SuplierController) ShowData(c *gin.Context) {
	var existingbarang models.SuplierModel
	id := c.Param("id")
	Fid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "not valid id",
		})
		return
	}

	if efrr := c.ShouldBindJSON(&existingbarang).Error; efrr != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "not valid id",
		})
		return
	}

	//find data
	if errd := ptk.repo.DB.First(&id, Fid).Scan(&existingbarang).Error; errd != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not valid id",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    existingbarang,
		"code":    200,
		"message": "not valid id",
	})

}

func (ptk *SuplierController) DeleteData(c *gin.Context) {
	var existingbarang models.SuplierModel
	id := c.Param("id")
	Fid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "not valid id",
		})
		return
	}

	if efrr := c.ShouldBindJSON(&existingbarang).Error; efrr != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "not valid id",
		})
		return
	}
	//find data
	if errd := ptk.repo.DB.First(&id, Fid).Scan(&existingbarang).Error; errd != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not valid id",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    existingbarang,
		"code":    200,
		"message": "not valid id",
	})

}
