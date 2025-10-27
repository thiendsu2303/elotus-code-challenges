package requests

type LoginRequest struct {
    // Username is the account identifier
    // example: alice
    Username string `json:"username" binding:"required"`
    // Password is the account secret
    // example: secret123
    Password string `json:"password" binding:"required"`
}
