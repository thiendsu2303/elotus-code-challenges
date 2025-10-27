package handler

import (
    "io"
    "net/http"
    "strings"

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

// UploadImage uploads an image file (max 8MB) and stores metadata
// @Summary Upload image
// @Description Upload an image file (PNG/JPEG/WebP) up to 8MB, store metadata
// @Tags resource
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param file formData file true "Image file to upload"
// @Success 201 {object} response.ImageResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 413 {object} response.ErrorResponse
// @Failure 415 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/v1/upload [post]
func (h *ResourceHandler) UploadImage(c *gin.Context) {
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

    // Read file from multipart form (prefer 'file', fallback to 'data')
    fileHeader, err := c.FormFile("file")
    if err != nil {
        fileHeader, err = c.FormFile("data")
        if err != nil {
            c.JSON(http.StatusBadRequest, response.NewError("missing file 'file'"))
            return
        }
    }

    src, err := fileHeader.Open()
    if err != nil {
        c.JSON(http.StatusBadRequest, response.NewError("cannot open uploaded file"))
        return
    }
    defer src.Close()

    // Detect content type from the first 512 bytes (validate type before size)
    sniff := make([]byte, 512)
    n, _ := src.Read(sniff)
    contentType := http.DetectContentType(sniff[:n])
    if !strings.HasPrefix(contentType, "image/") {
        c.JSON(http.StatusUnsupportedMediaType, response.NewError("only image content-type is allowed"))
        return
    }
    if seeker, ok := src.(io.Seeker); ok {
        _, _ = seeker.Seek(0, io.SeekStart)
    }

    // Check size limit after confirming it's an image
    const maxSize = 8 * 1024 * 1024 // 8MB
    if fileHeader.Size > maxSize {
        c.JSON(http.StatusRequestEntityTooLarge, response.NewError("file too large (max 8MB)"))
        return
    }

    userAgent := c.Request.UserAgent()
    clientIP := c.ClientIP()

    // Save file and persist metadata via service layer
    img, err := h.imageService.SaveUpload(uid, fileHeader, contentType, userAgent, clientIP)
    if err != nil {
        c.JSON(http.StatusInternalServerError, response.NewError("cannot save image metadata"))
        return
    }

    resp := response.ImageResponse{
        BaseResponse: response.BaseResponse{Status: "success", Message: "uploaded"},
        Data:         response.FromDomainImage(img),
    }
    c.JSON(http.StatusCreated, resp)
}