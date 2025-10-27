package response

type LoginData struct {
    // AccessToken is the JWT for API access
    // example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
    AccessToken string `json:"access_token"`
    // TokenType is typically "Bearer"
    // example: Bearer
    TokenType   string `json:"token_type"`
    // ExpiresAt is RFC3339 timestamp
    // example: 2025-01-01T12:00:00Z
    ExpiresAt   string `json:"expires_at"`
}

type LoginResponse struct {
    BaseResponse
    Data LoginData `json:"data"`
}
