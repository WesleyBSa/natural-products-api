package routes

import (
	"natural-products-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	productRoutes := router.Group("/products")
	{
		productRoutes.GET("/", controllers.GetProducts)
		productRoutes.POST("/", controllers.CreateProduct)
	}

	return router
}
