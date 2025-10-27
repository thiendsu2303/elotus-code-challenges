package main

import (
    "flag"
    "fmt"
    "log"
    "strings"
    "time"

    "backend-hackathon/internal/config"
    "backend-hackathon/internal/domain"
    "backend-hackathon/internal/repository"
)

func main() {
    var (
        revokeAll    bool
        revokeUsers  string
    )

    flag.BoolVar(&revokeAll, "all", false, "Revoke tokens for all users")
    flag.StringVar(&revokeUsers, "users", "", "Comma-separated usernames to revoke (e.g., alice,bob)")
    flag.Parse()

    if !revokeAll && revokeUsers == "" {
        fmt.Println("Usage: go run cmd/admin/main.go --all | --users alice,bob")
        flag.PrintDefaults()
        return
    }

    // Load config and init DB
    cfg := config.LoadConfig()
    db, err := config.InitDatabase(cfg)
    if err != nil {
        log.Fatalf("Failed to init database: %v", err)
    }

    userRepo := repository.NewUserRepository(db)
    now := time.Now().UTC()

    if revokeAll {
        res := db.Model(&domain.User{}).Update("revoked_at", now)
        if res.Error != nil {
            log.Fatalf("Failed to revoke all users: %v", res.Error)
        }
        fmt.Printf("Revoked tokens for all users at %s (rows affected: %d)\n", now.Format(time.RFC3339), res.RowsAffected)
        return
    }

    // Revoke specific users
    usernames := strings.Split(revokeUsers, ",")
    var success, failed int
    for _, u := range usernames {
        username := strings.TrimSpace(u)
        if username == "" {
            continue
        }
        user, err := userRepo.GetByUsername(username)
        if err != nil || user == nil {
            log.Printf("User not found: %s", username)
            failed++
            continue
        }
        user.RevokedAt = &now
        if err := userRepo.Update(user); err != nil {
            log.Printf("Failed to revoke user %s: %v", username, err)
            failed++
            continue
        }
        success++
    }

    fmt.Printf("Revoked tokens at %s. Success: %d, Failed: %d\n", now.Format(time.RFC3339), success, failed)
}