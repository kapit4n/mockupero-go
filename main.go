package main

import (
	"os"
	"strconv"

	db "github.com/kapit4n/go-mockupero/db"
	server "github.com/kapit4n/go-mockupero/server"
)

// main ...
func main() {
	database := db.Connect()
	s := server.Setup(database)
	port := "8080"

	if p := os.Getenv("PORT"); p != "" {
		if _, err := strconv.Atoi(p); err == nil {
			port = p
		}
	}

	s.Run(":" + port)
}
