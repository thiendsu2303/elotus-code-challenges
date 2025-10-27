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

// Ping godoc
// @Summary Health check
// @Description Returns pong
// @Tags system
// @Produce json
// @Success 200 {object} response.BaseResponse
// @Router /ping [get]
func (h *PingHandler) Ping(c *gin.Context) {
    c.JSON(http.StatusOK, response.BaseResponse{
        Status:  "success",
        Message: "pong",
    })
}

// PingWithAuth requires authentication and returns a success pong
// @Summary Authenticated ping
// @Description Requires Bearer token; returns pong_auth
// @Tags system
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.BaseResponse
// @Failure 401 {object} response.ErrorResponse
// @Router /api/v1/ping-auth [get]
func (h *PingHandler) PingWithAuth(c *gin.Context) {
    c.JSON(http.StatusOK, response.BaseResponse{
        Status:  "success",
        Message: "pong_auth",
    })
}
