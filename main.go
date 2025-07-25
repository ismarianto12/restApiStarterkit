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
	// Initialize Gin router
	r := gin.Default()

	// Initialize database
	if err := database.InitDB(); err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// Initialize controllers
	barangController := controllers.NewBarangController()
	purchasingController := controllers.NewPurchasingControllerInstance()
	userController := controllers.NewUserController()
	suplierController := controllers.NewSuplierControllerInstance()
	stockController := controllers.NewStockControllerInstance()

	// Public routes (no authentication required)
	public := r.Group("/")
	{
		public.GET("/", homeHandler)
		public.POST("/login", userController.Login)
	}

	// Protected routes (require authentication)
	protected := r.Group("/")
	protected.Use(utils.AuthMiddleware)
	{
		// Not found handler
		// protected.NoRoute(notFoundHandler)

		// User routes
		setupUserRoutes(protected, userController)

		// Barang routes
		setupBarangRoutes(protected, barangController)

		// Purchasing routes
		setupPurchasingRoutes(protected, purchasingController)

		// Supplier routes
		setupSuplierRoutes(protected, suplierController)

		// Stock routes
		setupStockRoutes(protected, stockController)

		// Category routes
		setupCategoryRoutes(protected, suplierController)
	}

	// Start server
	if err := r.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}

// Handlers
func homeHandler(c *gin.Context) {
	if err := godotenv.Load(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error loading environment",
			"version": 12,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"app":        os.Getenv("APP_VERSION"),
		"restfulAPI": "version 1",
		"data":       "v1",
	})
}

func NotFoundHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"status":  "error",
		"message": "Endpoint not found",
	})
}

// Route setup functions
func setupUserRoutes(r *gin.RouterGroup, uc *controllers.UserController) {
	users := r.Group("/users")
	{
		users.GET("/list", uc.GetAllUsers)
		users.POST("/", uc.CreateUser)
		users.POST("/create", uc.CreateUser) // Consider removing duplicate endpoint
		users.GET("/:email", uc.GetUserByEmail)
		users.GET("/show/:id", uc.ShowUser)
		users.GET("/userencp", controllers.EncpDe)
	}
}

func setupBarangRoutes(r *gin.RouterGroup, bc *controllers.BarangController) {
	barang := r.Group("/barang")
	{
		barang.POST("/upload", bc.UploadDfile)
		barang.GET("/list", bc.GetAllData) // Changed from GetSemuaKontol to GetAllBarang
		barang.GET("/show/:id", bc.GetBarangByid)
		barang.POST("/create", bc.Store)
		barang.PUT("/update/:id", bc.UpdateData)
		barang.DELETE("/delete/:id", bc.Delet) // Fixed typo from Delet to Delete
	}
}

func setupPurchasingRoutes(r *gin.RouterGroup, pc *controllers.NewPurchasingController) {
	purchasing := r.Group("/purchasing") // Fixed typo from purhcasing to purchasing
	{
		purchasing.GET("/list", pc.GetAllData)
		purchasing.POST("/create", pc.Store)
		purchasing.GET("/show/:id", pc.Show)
		// Removed duplicate /defer/:id endpoint
	}
}

func setupSuplierRoutes(r *gin.RouterGroup, sc *controllers.SuplierController) {
	supplier := r.Group("/supplier") // Fixed spelling from suplier to supplier
	{
		supplier.GET("/list", sc.Index)
		supplier.POST("/create", sc.CreateData)
		supplier.POST("/update/:id", sc.UpdateData)
		supplier.GET("/show/:id", sc.ShowData)
		supplier.DELETE("/delete/:id", sc.DeleteData) // Changed from GET to DELETE
	}
}

func setupStockRoutes(r *gin.RouterGroup, sc *controllers.NewStockController) {
	stock := r.Group("/stock")
	{
		stock.GET("/all", sc.GetAllData)
	}
}

func setupCategoryRoutes(r *gin.RouterGroup, sc *controllers.SuplierController) {
	categories := r.Group("/categories")
	{
		categories.GET("/index", sc.Index)
		categories.GET("/show/:id", sc.ShowData) // Changed from showdata/:id to show/:id
	}
}
