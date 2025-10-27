package config

import (
	"os"
)

type Config struct {
    ServerPort     string
    DBHost         string
    DBPort         string
    DBUser         string
    DBPassword     string
    DBName         string
    DBSSLMode      string
    JWTSecret      string
    JWTIssuer      string
    AccessTokenTTL string
}

func LoadConfig() *Config {
    return &Config{
        ServerPort:     getEnv("SERVER_PORT", "8080"),
        DBHost:         getEnv("DB_HOST", "localhost"),
        DBPort:         getEnv("DB_PORT", "5432"),
        DBUser:         getEnv("DB_USER", "postgres"),
        DBPassword:     getEnv("DB_PASSWORD", "postgres"),
        DBName:         getEnv("DB_NAME", "hackathon_db"),
        DBSSLMode:      getEnv("DB_SSL_MODE", "disable"),
        JWTSecret:      getEnv("JWT_SECRET", "dev-secret-change-me"),
        JWTIssuer:      getEnv("JWT_ISSUER", "backend-hackathon"),
        AccessTokenTTL: getEnv("ACCESS_TOKEN_TTL", "3600"),
    }
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func (c *Config) GetDSN() string {
	return "host=" + c.DBHost +
		" port=" + c.DBPort +
		" user=" + c.DBUser +
		" password=" + c.DBPassword +
		" dbname=" + c.DBName +
		" sslmode=" + c.DBSSLMode
}
