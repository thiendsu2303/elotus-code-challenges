package middleware

import (
    "fmt"
    "net/http"
    "strings"
    "time"

    "backend-hackathon/internal/repository"
    "backend-hackathon/internal/response"

    jwt "github.com/golang-jwt/jwt/v5"
    "github.com/gin-gonic/gin"
)

// AuthMiddleware verifies JWT access token and attaches user to context.
// It enforces issuer, expiry, and time-based revocation via users.revoked_at.
func AuthMiddleware(userRepo repository.UserRepository, jwtSecret string, jwtIssuer string) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewError("missing or invalid Authorization header"))
            return
        }

        tokenStr := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
        claims := jwt.MapClaims{}
        token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return []byte(jwtSecret), nil
        })
        if err != nil || !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewError("invalid token"))
            return
        }

        // Verify issuer
        if iss, ok := claims["iss"].(string); !ok || iss != jwtIssuer {
            c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewError("invalid token issuer"))
            return
        }

        // Verify expiry
        now := time.Now().UTC()
        var expUnix int64
        switch v := claims["exp"].(type) {
        case float64:
            expUnix = int64(v)
        case int64:
            expUnix = v
        case int:
            expUnix = int64(v)
        default:
            c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewError("invalid token expiry"))
            return
        }
        if time.Unix(expUnix, 0).Before(now) {
            c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewError("token expired"))
            return
        }

        // Extract subject (user id)
        var uid uint
        switch v := claims["sub"].(type) {
        case float64:
            uid = uint(v)
        case int64:
            uid = uint(v)
        case int:
            uid = uint(v)
        case string:
            // best-effort parse numeric id from string
            // if fails, treat as invalid subject
            c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewError("invalid token subject"))
            return
        default:
            c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewError("invalid token subject"))
            return
        }

        user, err := userRepo.GetByID(uid)
        if err != nil || user == nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewError("invalid user"))
            return
        }

        // Time-based revocation check: iat <= users.revoked_at means token invalid
        var iatUnix int64
        switch v := claims["iat"].(type) {
        case float64:
            iatUnix = int64(v)
        case int64:
            iatUnix = v
        case int:
            iatUnix = int64(v)
        default:
            c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewError("invalid token issued-at"))
            return
        }
        if user.RevokedAt != nil && iatUnix <= user.RevokedAt.Unix() {
            c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewError("token revoked"))
            return
        }

        // Attach to context and continue
        c.Set("user", user)
        c.Set("user_id", user.ID)
        c.Next()
    }
}

// RequireAuth ensures that a user has been set by AuthMiddleware.
func RequireAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        if _, exists := c.Get("user"); !exists {
            c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewError("unauthorized"))
            return
        }
        c.Next()
    }
}
