package handler

import (
	"net/http"
	"time"

	"backend-hackathon/internal/requests"
	"backend-hackathon/internal/response"
	"backend-hackathon/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req requests.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewError("invalid payload"))
		return
	}

	user, err := h.authService.Register(req.Username, req.Password)
	if err != nil {
		if err.Error() == "username already exists" {
			c.JSON(http.StatusConflict, response.NewError("username already exists"))
			return
		}
		c.JSON(http.StatusInternalServerError, response.NewError("internal error"))
		return
	}

	resp := response.RegisterResponse{
		BaseResponse: response.BaseResponse{Status: "success", Message: "registered"},
		Data:         response.RegisterData{ID: user.ID, Username: user.Username, CreatedAt: user.CreatedAt},
	}
	c.JSON(http.StatusCreated, resp)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req requests.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewError("invalid payload"))
		return
	}

	token, exp, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.NewError("invalid credentials"))
		return
	}

	resp := response.LoginResponse{
		BaseResponse: response.BaseResponse{Status: "success", Message: "logged_in"},
		Data: response.LoginData{
			AccessToken: token,
			TokenType:   "Bearer",
			ExpiresAt:   exp.Format(time.RFC3339),
		},
	}
	c.JSON(http.StatusOK, resp)
}

func (h *AuthHandler) Logout(c *gin.Context) {
    uidAny, ok := c.Get("user_id")
    if !ok {
        c.JSON(http.StatusUnauthorized, response.NewError("unauthorized"))
        return
    }

    var uid uint
    switch v := uidAny.(type) {
    case uint:
        uid = v
    case int:
        uid = uint(v)
    case int64:
        uid = uint(v)
    case float64:
        uid = uint(v)
    default:
        c.JSON(http.StatusUnauthorized, response.NewError("unauthorized"))
        return
    }

    if err := h.authService.Logout(uid); err != nil {
        c.JSON(http.StatusInternalServerError, response.NewError("internal error"))
        return
    }

    c.JSON(http.StatusOK, response.BaseResponse{Status: "success", Message: "logged_out"})
}
