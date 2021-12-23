package server

import (
	"github.com/gin-contrib/cors"
	"github.com/kapit4n/go-mockupero/middleware"
	"github.com/kapit4n/go-mockupero/router"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))

	r.Use(middleware.SetDBtoContext(db))

	r.Use(cors.Default())

	router.Initialize(r)
	return r
}
