package db

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
    err := godotenv.Load("../.env")
    if err != nil {
        log.Printf("Could not load .env file: %v", err)
    }

    wd, err := os.Getwd()
    if err != nil {
        log.Fatalf("Could not get working directory: %v", err)
    }
    migrationPath := filepath.Join(wd, "migrations")

    os.Setenv("MIGRATION_PATH", "file://"+migrationPath)

    code := m.Run()
    os.Exit(code)
}