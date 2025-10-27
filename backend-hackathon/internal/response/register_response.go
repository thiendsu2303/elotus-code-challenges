package response

import "time"

type RegisterData struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterResponse struct {
	BaseResponse
	Data RegisterData `json:"data"`
}
