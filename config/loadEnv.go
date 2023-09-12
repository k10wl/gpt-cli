package config

import (
	"log"
	"path"

	"github.com/joho/godotenv"
)

func loadEnv(dir string) {
	file := path.Join(dir, ".env")
	err := godotenv.Load(file)
	if err != nil {
		log.Fatalf("Error upon loading env: %+v", err)
	}
}
