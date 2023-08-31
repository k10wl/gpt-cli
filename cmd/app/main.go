package main

import (
	"log"
	"os"
	"path/filepath"

	input "cli/internal/input"

	"github.com/joho/godotenv"
)

func main() {
	exePath, err := os.Executable()
	exeDir := filepath.Dir(exePath)
	envPath := filepath.Join(exeDir, ".env")
	err = godotenv.Load(envPath)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	input.Start()
}
