/*
* Golang REST API Example
* This is a simple REST API example using Golang and Gin framework.
* It includes a controller for managing "barang" (items) and a user controller.
* @author Ismarianto
* @email : ismarianto97@gmail.com
**/

package main

import (
	"golangRest/src/controllers"
	"golangRest/src/database"
	"golangRest/src/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	err := database.InitDB()
	if err != nil {
		panic(err)
	}
	r.Use(utils.InterCeptor())
	barangController := controllers.NewBarangController()
	purchasingController := controllers.NewPurchasingControllerInstance()
	userController := controllers.NewUserController()
	suplierController := controllers.NewSuplierControllerInstance()

	r.GET("/", func(c *gin.Context) {
		err := godotenv.Load()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error page", "version": 12})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"app":         os.Getenv("APP_VERSION"),
			"restfullAPi": "version 1",
			"data":        "v1",
		})
	})
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"page":     "not found",
			"response": "ada",
		})
	})
	r.POST("/login", userController.Login)
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/list", userController.GetAllUsers)
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.POST("/create", userController.CreateUser)
		userRoutes.GET("/:email", userController.GetUserByEmail)
		userRoutes.GET("/show/:id", userController.ShowUser)
	}

	barangRoutes := r.Group("/barang")
	{
		barangRoutes.POST("/upload", barangController.UploadDfile)
		barangRoutes.GET("/list", barangController.GetSemuaKontol)
		barangRoutes.GET("/show/:id", barangController.GetBarangByid)

		barangRoutes.POST("/create", barangController.Store)
		barangRoutes.PUT("/update/:id", barangController.UpdateData)
		barangRoutes.DELETE("/delete/:id", barangController.Delet)
	}
	purcahsing := r.Group("/purhcasing")
	{
		purcahsing.GET("/list", purchasingController.GetAllData)
		purcahsing.POST("/create", purchasingController.Store)
		purcahsing.GET("/show/:id", purchasingController.Show)

	}
	suplier := r.Group("/suplier")
	{
		suplier.GET("/list", suplierController.Index)
		suplier.POST("/create", suplierController.CreateData)
		suplier.POST("/update/:id", suplierController.UpdateData)
		suplier.GET("/show/:id", suplierController.ShowData)
		suplier.GET("/delete/:id", suplierController.DeleteData)

	}

	r.Run(":8080")
}
