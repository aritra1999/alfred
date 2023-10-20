package main

import (
	"albert/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	public := router.Group("/api")

	public.POST("/signup", controllers.SignUp)

	router.Run(":8080")

}
