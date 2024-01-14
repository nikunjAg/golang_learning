package router

import (
	"example.com/events-app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRoutes(server *gin.Engine) {

	// Authenticated Group
	authenticatedGroup := server.Group("/")
	authenticatedGroup.Use(middlewares.ValidateToken)

	// Events routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)

	// AuthRequired Event Routes
	authenticatedGroup.POST("/events", createEvent)
	authenticatedGroup.PUT("/events/:id", updateEventById)
	authenticatedGroup.DELETE("/events/:id", deleteEventById)

	// Users routes
	server.POST("/users/signup", createUser)
	server.POST("/users/login", loginUser)
}
