package requests

type RegisterRequest struct {
    // Username is the desired login name (3-50 chars)
    // example: alice
    Username string `json:"username" binding:"required,min=3,max=50"`
    // Password must be at least 6 characters
    // example: secret123
    Password string `json:"password" binding:"required,min=6"`
}
