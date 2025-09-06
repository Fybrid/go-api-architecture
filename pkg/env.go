package pkg

import (
    "fmt"

    "github.com/joho/godotenv"
)

func LoadDotenv(paths []string) error {
    for _, p := range paths {
        if err := godotenv.Load(p); err == nil {
            return nil
        }
    }
    return fmt.Errorf("could not load .env from any of: %v", paths)
}

func LoadEnv() {
    _ = LoadDotenv([]string{".env", "../.env", "../../.env"})
}
