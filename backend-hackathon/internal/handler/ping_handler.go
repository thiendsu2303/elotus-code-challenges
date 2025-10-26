package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingHandler handles ping requests
type PingHandler struct{}

// NewPingHandler creates a new ping handler
func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

// Ping handles ping request
func (h *PingHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"status":  "ok",
	})
}
