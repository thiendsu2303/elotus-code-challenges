package handler

import (
    "net/http"

    "backend-hackathon/internal/response"
    "backend-hackathon/internal/service"

    "github.com/gin-gonic/gin"
)

type ResourceHandler struct {
    imageService service.ImageService
}

func NewResourceHandler(imageService service.ImageService) *ResourceHandler {
    return &ResourceHandler{imageService: imageService}
}

// ListMyImages returns images belonging to the authenticated user (from token)
// @Summary List my images
// @Description Returns images belonging to the authenticated user
// @Tags resource
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.ImagesResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/v1/resource/images [get]
func (h *ResourceHandler) ListMyImages(c *gin.Context) {
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

    images, err := h.imageService.ListMyImages(uid)
    if err != nil {
        c.JSON(http.StatusInternalServerError, response.NewError("internal error"))
        return
    }

    resp := response.ImagesResponse{
        BaseResponse: response.BaseResponse{Status: "success", Message: "images"},
        Data:         response.FromDomainImages(images),
    }
    c.JSON(http.StatusOK, resp)
}