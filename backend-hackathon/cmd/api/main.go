package main

import (
    "log"

    "backend-hackathon/internal/config"
    "backend-hackathon/internal/handler"
    "backend-hackathon/internal/repository"
    "backend-hackathon/internal/service"
    "backend-hackathon/internal/router"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := config.InitDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

    // Initialize layers
    userRepo := repository.NewUserRepository(db)
    authService := service.NewAuthService(userRepo)
    authHandler := handler.NewAuthHandler(authService)
    pingHandler := handler.NewPingHandler()

	// Setup router
	r := router.SetupRouter(pingHandler, authHandler)

	// Start server
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
