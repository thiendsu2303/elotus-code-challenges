package response

type LoginData struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresAt   string `json:"expires_at"`
}

type LoginResponse struct {
	BaseResponse
	Data LoginData `json:"data"`
}
