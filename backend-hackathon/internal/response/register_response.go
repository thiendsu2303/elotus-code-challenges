package response

import "time"

type RegisterData struct {
    // ID is the user identifier
    // example: 1
    ID        uint      `json:"id"`
    // Username is the registered username
    // example: alice
    Username  string    `json:"username"`
    // CreatedAt is the account creation time (RFC3339)
    // example: 2025-01-01T12:00:00Z
    CreatedAt time.Time `json:"created_at"`
}

type RegisterResponse struct {
    BaseResponse
    Data RegisterData `json:"data"`
}
