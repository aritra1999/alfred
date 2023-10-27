package main

import (
	"albert/controllers"
	"albert/models"

	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDataBase()
	router := gin.Default()
	public := router.Group("/api")

	authRouter := public.Group("/auth")
	{
		authRouter.POST("/signup", controllers.SignUp)
	}

	router.Run(":8080")

}
