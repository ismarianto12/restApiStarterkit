/**
*@Author : ismarianto
*
*
 **/

package controllers

import (
	"golangRest/src/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	repo *repository.CategoryRepository
}

func NewCategoryRepository(rp *repository.CategoryRepository) *CategoryController {
	return &CategoryController{repo: rp}
}

func (crot *CategoryController) Index(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"messages": "data"})
}

func (crot *CategoryController) Show(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"messages": "data"})
}

func (crot *CategoryController) Update(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"messages": "data"})
}

func (crot *CategoryController) Delete(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"messages": "data"})
}
