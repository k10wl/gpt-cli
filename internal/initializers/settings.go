package initializers

import (
	"errors"
	"strings"
)

type Settings struct {
	Model     string
	Assistant string
}

func loadSettings(config *Config, flags *Flags) (*Settings, error) {
	model, err := getModel(config, flags)
	if err != nil {
		return &Settings{}, err
	}

	system, err := getAssistantMessage(config, flags)
	if err != nil {
		return &Settings{}, err
	}

	return &Settings{
		Model:     *model,
		Assistant: system,
	}, nil
}

func getModel(config *Config, flags *Flags) (*string, error) {
	model := flags.Model

	if model == "" {
		return &config.Models[0], nil
	}

	exists := false
	for _, m := range config.Models {
		if m == model {
			exists = true
			break
		}
	}

	if !exists {
		errorMessage := "Requested model (" + model + ") is not supported. Available: " + strings.Join(config.Models, ", ")
		return nil, errors.New(errorMessage)
	}

	return &model, nil
}

func getAssistantMessage(config *Config, flags *Flags) (string, error) {
	system := flags.Assistant

	if system == "" {
		return "", nil
	}

	message, exist := config.Assistant[system]
	if !exist {
		errorMessage := "Requested system message (" + message + ") is not supported. Available: " + listAvailableAssistant(config)
		return "", errors.New(errorMessage)
	}

	return message, nil
}

func listAvailableAssistant(config *Config) string {
	available := make([]string, 0, len(config.Assistant))
	for k := range config.Assistant {
		available = append(available, k)
	}

	return strings.Join(available, ", ")
}
