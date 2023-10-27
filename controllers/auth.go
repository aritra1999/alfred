package controllers

import (
	service "albert/services"
	. "albert/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {

	var input SignUpInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, service.SignUp(input))
}
