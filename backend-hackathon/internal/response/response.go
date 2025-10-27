package response

type BaseResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}
