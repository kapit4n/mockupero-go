package server

import (
	"github.com/kapit4n/go-mockupero/middleware"
	"github.com/kapit4n/go-mockupero/router"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.SetDBtoContext(db))
	router.Initialize(r)
	return r
}
