package initializers

import (
	"path"

	"github.com/joho/godotenv"
)

func LoadEnv(flags *Flags) error {
	file := path.Join(flags.Path, ".env")
	return godotenv.Load(file)
}
