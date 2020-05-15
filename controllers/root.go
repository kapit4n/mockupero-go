package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func APIEndpoints(c *gin.Context) {
	reqScheme := "http"

	if c.Request.TLS != nil {
		reqScheme = "https"
	}

	reqHost := c.Request.Host
	baseURL := fmt.Sprintf("%s://%s", reqScheme, reqHost)

	resources := map[string]string{
		"features_url": baseURL + "/features",
		"feature_url":  baseURL + "/features/{id}",
		"mockups_url":  baseURL + "/mockups",
		"mockup_url":   baseURL + "/mockups/{id}",
		"projects_url": baseURL + "/projects",
		"project_url":  baseURL + "/projects/{id}",
		"users_url":    baseURL + "/users",
		"user_url":     baseURL + "/users/{id}",
		"comments_url": baseURL + "/comments",
		"comment_url":  baseURL + "/comments/{id}",
	}

	c.IndentedJSON(http.StatusOK, resources)
}
