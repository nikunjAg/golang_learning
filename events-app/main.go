package main

import (
	"fmt"

	"example.com/events-app/db"
	"example.com/events-app/router"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	_, err := db.InitDB()

	if err != nil {
		fmt.Println(err)
		return
	}

	router.InitRoutes(server)

	server.Run(":8080") // localhost:8080
}
