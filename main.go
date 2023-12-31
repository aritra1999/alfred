package main

import (
	"albert/controllers"
	_ "albert/docs"
	"albert/middlewares"
	"albert/models"
	"albert/services/cache"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/time/rate"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
	models.ConnectDataBase()
	models.MigrateTables()
	cache.RedisClient()
}

// @title           CareerLab - Alfred
// @version         1.0
// @description     Alfred is our rest endpoint service
// @host      		localhost:8080
// @BasePath  		/v1/
func main() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// Middlewares
	router.Use(middlewares.CorsMiddleware())
	router.Use(middlewares.RateLimiterMiddleware(rate.Every(1*time.Second), 5))

	// Internal routes
	router.GET("/ping", controllers.Ping)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Public routes
	v1Public := router.Group("/v1")
	authRouter := v1Public.Group("/auth")

	authRouter.POST("/signup", controllers.SignUp)
	authRouter.POST("/send-magic-link", controllers.SendMagicLink)
	authRouter.GET("/magic-link", controllers.ValidateMagicLink)

	// Private routes
	v1Private := router.Group("/v1")
	v1Private.Use(middlewares.JwtAuthMiddleware())

	// Admin Routes
	adminRouter := router.Group("/v1")
	adminRouter.Use(middlewares.AdminAuthMiddleware())

	port := ":" + os.Getenv("PORT")
	router.Run(port)
}
