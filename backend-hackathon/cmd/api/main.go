package main

import (
	"log"

	"backend-hackathon/internal/config"
	"backend-hackathon/internal/handler"
	"backend-hackathon/internal/repository"
	"backend-hackathon/internal/router"
	"backend-hackathon/internal/usecase"
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
	registerUsecase := usecase.NewRegisterUsecase(userRepo)
	registerHandler := handler.NewRegisterHandler(registerUsecase)
	pingHandler := handler.NewPingHandler()

	// Setup router
	r := router.SetupRouter(pingHandler, registerHandler)

	// Start server
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
