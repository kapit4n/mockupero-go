package db

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/kapit4n/go-mockupero/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/serenize/snaker"
)

func Connect() *gorm.DB {
	dir := filepath.Dir("db/database3.db")
	db, err := gorm.Open("sqlite3", dir+"/database3.db")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}

	db.LogMode(false)

	if gin.IsDebugging() {
		db.LogMode(true)
	}

	fmt.Println("HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHhhh")

	if os.Getenv("AUTOMIGRATE") != "1" {
		db.AutoMigrate(
			&models.Feature{},
			&models.Mockup{},
			&models.MockupItem{},
			&models.Project{},
			&models.User{},
			&models.Comment{},
		)
	}

	return db
}

func DBInstance(c *gin.Context) *gorm.DB {
	return c.MustGet("DB").(*gorm.DB)
}

func (self *Parameter) SetPreloads(db *gorm.DB) *gorm.DB {
	if self.Preloads == "" {
		return db
	}

	for _, preload := range strings.Split(self.Preloads, ",") {
		var a []string

		for _, s := range strings.Split(preload, ".") {
			a = append(a, snaker.SnakeToCamel(s))
		}

		db = db.Preload(strings.Join(a, "."))
	}

	return db
}
