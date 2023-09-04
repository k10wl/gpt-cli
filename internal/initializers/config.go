package initializers

import (
	"encoding/json"
	"os"
	"path"
)

type Config struct {
	Models    []string          `json:"models"`
	Assistant map[string]string `json:"assistant"`
}

func readConfigFile(flags *Flags) (*Config, error) {
	configPath := path.Join(flags.Path, "config.json")

	byte, err := os.ReadFile(configPath)
	if err != nil {
		return &Config{}, err
	}

	var config Config
	json.Unmarshal(byte, &config)

	return &config, nil
}
