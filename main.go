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

	public.POST("/signup", controllers.SignUp)

	router.Run(":8080")

}
