package router

import (
	"backend-hackathon/internal/handler"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures all routes and returns a configured Gin engine
func SetupRouter(
	pingHandler *handler.PingHandler,
	registerHandler *handler.RegisterHandler,
) *gin.Engine {
	r := gin.Default()

	// Health check endpoint
	r.GET("/ping", pingHandler.Ping)

	// API routes
	api := r.Group("/api/v1")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", registerHandler.Register)
		}
	}

	return r
}
