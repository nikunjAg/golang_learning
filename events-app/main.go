package main

import (
	"fmt"
	"net/http"

	"example.com/events-app/db"
	"example.com/events-app/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {

	events, err := models.GetAllEvents()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get all events",
			"errors":  []string{err.Error()},
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Fetched all events successfully",
		"data": gin.H{
			"events": events,
		},
	})
}

func createEvent(context *gin.Context) {

	var event *models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	event.UserId = 2
	err = event.Save()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"message": "Event cretaed successfully",
		"event":   event,
	})
}

func main() {
	server := gin.Default()
	_, err := db.InitDB()

	if err != nil {
		fmt.Println(err)
		return
	}

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080") // localhost:8080
}
