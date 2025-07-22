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

	"github.com/gin-gonic/gin"
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

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/list", userController.GetAllUsers)
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.GET("/:email", userController.GetUserByEmail)
	}

	barangRoutes := r.Group("/barang")
	{
		barangRoutes.POST("/upload", barangController.UploadDfile)
		barangRoutes.GET("/list", barangController.GetSemuaKontol)
		barangRoutes.POST("/create", barangController.Store)
		barangRoutes.PUT("/update/:id", barangController.UpdateData)
	}
	purcahsing := r.Group("/purhcasing")
	{
		purcahsing.GET("/list", purchasingController.GetAllData)
		purcahsing.POST("/create", purchasingController.GetAllData)
	}

	r.Run(":8080")
}
