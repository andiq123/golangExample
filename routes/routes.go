package routes

import (
	"api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticatedGroup := server.Group("/")
	authenticatedGroup.Use(middlewares.Authenticate)
	authenticatedGroup.POST("/events", createEvent)
	authenticatedGroup.PUT("/events/:id", updateEvent)
	authenticatedGroup.DELETE("/events/:id", deleteEvent)
	authenticatedGroup.POST("/events/:id/register", registerForEvent)

	server.POST("/signup", signUp)
	server.POST("/login", signIn)
}
