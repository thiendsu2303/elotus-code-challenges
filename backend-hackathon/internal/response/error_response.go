package response

type ErrorResponse struct {
	BaseResponse
}

func NewError(message string) ErrorResponse {
	return ErrorResponse{
		BaseResponse: BaseResponse{Status: "error", Message: message},
	}
}
