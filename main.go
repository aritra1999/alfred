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
// @description     Alfred is our rest endpoint service
// @host      		localhost:8080
// @BasePath  		/api/
func main() {

	models.ConnectDataBase()

	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

	public := router.Group("/api")

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	authRouter := public.Group("/auth")
	{
		authRouter.POST("/signup", controllers.SignUp)
	}

	router.Run(":8080")

}
