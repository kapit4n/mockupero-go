package router

import (
	"github.com/kapit4n/go-mockupero/controllers"

	"github.com/gin-gonic/gin"
)

func Initialize(r *gin.Engine) {
	r.GET("/", controllers.APIEndpoints)

	api := r.Group("")
	{

		api.GET("/features", controllers.GetFeatures)
		api.GET("/features/:id", controllers.GetFeature)
		api.POST("/features", controllers.CreateFeature)
		api.PUT("/features/:id", controllers.UpdateFeature)
		api.DELETE("/features/:id", controllers.DeleteFeature)

		api.GET("/mockups", controllers.GetMockups)
		api.GET("/mockups/:id", controllers.GetMockup)
		api.POST("/mockups", controllers.CreateMockup)
		api.PUT("/mockups/:id", controllers.UpdateMockup)
		api.DELETE("/mockups/:id", controllers.DeleteMockup)

		api.GET("/projects", controllers.GetProjects)
		api.GET("/projects/:id", controllers.GetProject)
		api.POST("/projects", controllers.CreateProject)
		api.PUT("/projects/:id", controllers.UpdateProject)
		api.DELETE("/projects/:id", controllers.DeleteProject)

		api.GET("/users", controllers.GetUsers)
		api.GET("/users/:id", controllers.GetUser)
		api.POST("/users", controllers.CreateUser)
		api.PUT("/users/:id", controllers.UpdateUser)
		api.DELETE("/users/:id", controllers.DeleteUser)

	}
}
