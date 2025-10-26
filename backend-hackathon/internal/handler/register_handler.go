package handler

import (
	"net/http"

	"backend-hackathon/internal/requests"
	"backend-hackathon/internal/usecase"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	registerUsecase usecase.RegisterUsecase
}

func NewRegisterHandler(registerUsecase usecase.RegisterUsecase) *RegisterHandler {
	return &RegisterHandler{registerUsecase: registerUsecase}
}

func (h *RegisterHandler) Register(c *gin.Context) {
	var req requests.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.registerUsecase.Register(req.Username, req.Password)
	if err != nil {
		if err.Error() == "username already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": "username already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
