package controllers

import (
	"albert/models"
	"albert/services/auth"
	"albert/services/cache"
	"albert/services/mail"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type SignUpInput struct {
	Email string `json:"email" binding:"required"`
}

func SignUp(c *gin.Context) {
	var input SignUpInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}
	u.Email = input.Email

	if _, err := u.SaveUser(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mail.SendEmail(u.Email, "Welcome to CareerLab", "You have successfully signed up to CareerLab")

	c.JSON(http.StatusOK, gin.H{"message": "user created successfully"})
}

type MagicLinkInput struct {
	Email string `json:"email" binding:"required"`
}

func SendMagicLink(c *gin.Context) {

	var input MagicLinkInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Email: input.Email,
	}

	if !user.Exists() {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	magicLinkToken := auth.GenerateMagicLinkToken()
	cache.Set(magicLinkToken, user.Email, 10*time.Minute)
	frontEnd := os.Getenv("FRONTEND_URL")
	magicLink := fmt.Sprintf("%s/auth/validate?token=%s", frontEnd, magicLinkToken)

	emailBody := fmt.Sprintf("Your <a href='%s' target='_blank'>magic link</a> is here! (valid for 10 mins)", magicLink)

	if _, err := mail.SendEmail(user.Email, "CareerLab OTP", emailBody); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending magic link, please try again later"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Magic link sent to %s, please check email", user.Email)})
}

func ValidateMagicLink(c *gin.Context) {
	token := c.Query("token")

	email, err := cache.Get(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	user, err := models.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	jwtToken, err := user.GenerateToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}
