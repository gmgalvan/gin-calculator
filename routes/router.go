package routes

import (
	docs "github.com/gin-calculator/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New roiter returns a gin engine router
func NewRouter() *gin.Engine {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/calculator/v1"

	router.POST("/calculator/v1/sum", Sum)

	// use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
