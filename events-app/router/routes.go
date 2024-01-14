package router

import "github.com/gin-gonic/gin"

func InitRoutes(server *gin.Engine) {

	// Events routes
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.GET("/events/:id", getEventById)
	server.PUT("/events/:id", updateEventById)
	server.DELETE("/events/:id", deleteEventById)

	// Users routes
	server.POST("/users/signup", createUser)
	server.POST("/users/login", loginUser)
}
