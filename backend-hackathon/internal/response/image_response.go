package response

// ImageResponse wraps a single image item
type ImageResponse struct {
    BaseResponse
    Data ImageItem `json:"data"`
}