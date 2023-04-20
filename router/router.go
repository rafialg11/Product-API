package router

import (
	"product-api/controllers"
	"product-api/database"
	"product-api/middlewares"
	"product-api/repository"
	"product-api/service"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	db := database.GetDB()
	r := gin.Default()
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controllers.NewProductController(productService)
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		//user and admin can access
		productRouter.POST("/", middlewares.UserAuthorization(), controllers.CreateProduct)
		productRouter.GET("/:productId", middlewares.UserAuthorization(), middlewares.AccessByIdAuthorization(), productController.GetProductById)
		//only admin
		productRouter.GET("/", middlewares.AdminAuthorization(), productController.GetAllProduct)
		productRouter.DELETE("/:productId", middlewares.AdminAuthorization(), controllers.DeleteProduct)
		productRouter.PUT("/:productId", middlewares.AdminAuthorization(), controllers.UpdateProduct)
	}

	return r
}
