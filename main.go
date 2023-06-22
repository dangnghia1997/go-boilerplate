package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"main/config"
)

func main() {
	db, err := config.NewDB()
	if err != nil {
		log.Fatal("could not initialize database: %w", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
