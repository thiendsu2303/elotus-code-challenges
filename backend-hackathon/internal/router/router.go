package router

import (
	"backend-hackathon/internal/handler"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures all routes and returns a configured Gin engine
func SetupRouter(
	userHandler *handler.UserHandler,
) *gin.Engine {
	r := gin.Default()

	// Health check endpoint
	r.GET("/ping", userHandler.Ping)

	// API routes
	api := r.Group("/api/v1")
	{
		// User routes
		users := api.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("", userHandler.GetAllUsers)
			users.GET("/:id", userHandler.GetUserByID)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}
	}

	return r
}
