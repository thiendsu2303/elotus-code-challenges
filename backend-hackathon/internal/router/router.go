package router

import (
    "backend-hackathon/internal/handler"
    "time"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter configures all routes and returns a configured Gin engine
func SetupRouter(
    pingHandler *handler.PingHandler,
    authHandler *handler.AuthHandler,
    resourceHandler *handler.ResourceHandler,
    authMW gin.HandlerFunc,
) *gin.Engine {
    r := gin.Default()

    // Enable CORS for frontend dev server (localhost:3000)
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Authorization", "Content-Type"},
        ExposeHeaders:    []string{"Authorization", "Content-Type"},
        AllowCredentials: false,
        MaxAge:           12 * time.Hour,
    }))

	// Health check endpoint
	r.GET("/ping", pingHandler.Ping)

    // Swagger UI
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // API routes
    api := r.Group("/api/v1")
    {
        // Auth routes
        auth := api.Group("/auth")
        {
            auth.POST("/register", authHandler.Register)
            auth.POST("/login", authHandler.Login)
        }

        // Protected routes
        protected := api.Group("/")
        protected.Use(authMW)
        {
            protected.GET("/ping-auth", pingHandler.PingWithAuth)
            protected.POST("/auth/logout", authHandler.Logout)
            // Resource routes
            protected.GET("/resource/images", resourceHandler.ListMyImages)
        }
    }

    return r
}
