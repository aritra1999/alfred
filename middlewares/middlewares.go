package middlewares

import (
	"albert/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := utils.TokenValid(c); err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := utils.AdminValid(c); err != nil {
			c.String(http.StatusUnauthorized, "Only admins can access this route")
			c.Abort()
			return
		}
		c.Next()
	}
}
