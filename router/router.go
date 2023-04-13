package router

import (
	"product-api/controllers"
	"product-api/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

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
		productRouter.GET("/:productId", middlewares.UserAuthorization(), middlewares.AccessByIdAuthorization(), controllers.GetProductById)
		//only admin
		productRouter.GET("/", middlewares.AdminAuthorization(), controllers.GetAllProduct)
		productRouter.DELETE("/:productId", middlewares.AdminAuthorization(), controllers.DeleteProduct)
		productRouter.PUT("/:productId", middlewares.AdminAuthorization(), controllers.UpdateProduct)
	}

	return r
}
