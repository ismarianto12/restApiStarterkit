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

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	barangController := controllers.NewBarangController()
	userController := controllers.NewUserController()

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", userController.GetAllUsers)
	}

	barangRoutes := r.Group("/barang")
	{
		barangRoutes.GET("/list", barangController.GetAllData)
	}

	r.Run(":8080")
}
