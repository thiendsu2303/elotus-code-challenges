package handler

import (
	"net/http"

	"backend-hackathon/internal/response"

	"github.com/gin-gonic/gin"
)

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, response.BaseResponse{
		Status:  "success",
		Message: "pong",
	})
}
