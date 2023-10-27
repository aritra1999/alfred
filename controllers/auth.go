package controllers

import (
	service "albert/services"
	. "albert/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignUp      	 godoc
// @Summary      SignUp endpoint for creating new user
// @Description  Creates new user
// @Tags         auth
// @Produce      json
// @Param        user  body  SignUpInput  true  "User details"
// @Success      200  {array}  SignUpOutput
// @Router       /auth/signup [post]
func SignUp(c *gin.Context) {

	var input SignUpInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, service.SignUp(input))
}
