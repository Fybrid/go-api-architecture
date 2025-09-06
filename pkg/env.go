package pkg

import (
    "log"

    "github.com/joho/godotenv"
)

func LoadEnv() {
    // Try common locations so it works from repo root or cmd/app
    candidates := []string{".env", "../.env", "../../.env"}
    for _, p := range candidates {
        if err := godotenv.Load(p); err == nil {
            return
        }
    }
    log.Fatal("Error loading .env file")
}
