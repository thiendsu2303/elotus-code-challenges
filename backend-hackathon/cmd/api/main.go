package main

import (
    "log"
    "strconv"
    "time"

    "backend-hackathon/internal/config"
    "backend-hackathon/internal/handler"
    "backend-hackathon/internal/middleware"
    "backend-hackathon/internal/repository"
    "backend-hackathon/internal/router"
    "backend-hackathon/internal/service"
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
    imageRepo := repository.NewImageRepository(db)
    // Parse Access Token TTL (seconds)
    ttlSeconds, err := strconv.Atoi(cfg.AccessTokenTTL)
    if err != nil || ttlSeconds <= 0 {
        ttlSeconds = 3600
    }
    authService := service.NewAuthService(userRepo, cfg.JWTSecret, cfg.JWTIssuer, time.Duration(ttlSeconds)*time.Second)
    imageService := service.NewImageService(imageRepo)
    authHandler := handler.NewAuthHandler(authService)
    resourceHandler := handler.NewResourceHandler(imageService)
    pingHandler := handler.NewPingHandler()
    authMW := middleware.AuthMiddleware(userRepo, cfg.JWTSecret, cfg.JWTIssuer)

	// Setup router
    r := router.SetupRouter(pingHandler, authHandler, resourceHandler, authMW)

	// Start server
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
