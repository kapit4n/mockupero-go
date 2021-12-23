package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	session := sessions.Default(c)

	var loginRequest LoginRequest
	if err := c.BindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error to bind json"})
	}

	if strings.Trim(loginRequest.Email, " ") == "" || strings.Trim(loginRequest.Password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email or password should not be empty"})
		return
	}

	if loginRequest.Email != "test@gmail.com" || loginRequest.Password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfuly authenticated user"})
}
