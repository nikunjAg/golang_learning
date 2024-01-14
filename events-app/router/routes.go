package router

import "github.com/gin-gonic/gin"

func InitRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.GET("/events/:id", getEventById)
	server.PUT("/events/:id", updateEventById)
	server.DELETE("/events/:id", deleteEventById)
}
