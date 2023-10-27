package main

import (
	"albert/controllers"
	_ "albert/docs"
	"albert/models"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           CareerLab - Alfred
// @version         1.0
// @description     API documentation for Alfred

// @host      localhost:8080
// @BasePath  /api/

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	models.ConnectDataBase()

	router := gin.Default()
	public := router.Group("/api")

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	authRouter := public.Group("/auth")
	{
		authRouter.POST("/signup", controllers.SignUp)
	}

	router.Run(":8080")

}
