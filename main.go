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
	// r.Use(utils.InterCeptor())
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
		userRoutes.GET("/list", utils.AuthMiddleware, userController.GetAllUsers)
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.POST("/create", userController.CreateUser)
		userRoutes.GET("/:email", userController.GetUserByEmail)
		userRoutes.GET("/show/:id", userController.ShowUser)
	}

	barangRoutes := r.Group("/barang")
	{
		barangRoutes.POST("/upload", utils.AuthMiddleware, barangController.UploadDfile)
		barangRoutes.GET("/list", utils.AuthMiddleware, barangController.GetSemuaKontol)
		barangRoutes.GET("/show/:id", utils.AuthMiddleware, barangController.GetBarangByid)

		barangRoutes.POST("/create", utils.AuthMiddleware, barangController.Store)
		barangRoutes.PUT("/update/:id", utils.AuthMiddleware, barangController.UpdateData)
		barangRoutes.DELETE("/delete/:id", utils.AuthMiddleware, barangController.Delet)
	}
	purcahsing := r.Group("/purhcasing")
	{
		purcahsing.GET("/list", utils.AuthMiddleware, purchasingController.GetAllData)
		purcahsing.POST("/create", utils.AuthMiddleware, purchasingController.Store)
		purcahsing.GET("/show/:id", utils.AuthMiddleware, purchasingController.Show)

	}
	suplier := r.Group("/suplier")
	{
		suplier.GET("/list", utils.AuthMiddleware, suplierController.Index)
		suplier.POST("/create", utils.AuthMiddleware, suplierController.CreateData)
		suplier.POST("/update/:id", utils.AuthMiddleware, suplierController.UpdateData)
		suplier.GET("/show/:id", utils.AuthMiddleware, suplierController.ShowData)
		suplier.GET("/delete/:id", utils.AuthMiddleware, suplierController.DeleteData)
	}

	r.Run(":8080")
}
