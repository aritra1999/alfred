package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCountOfUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "world"})
}
