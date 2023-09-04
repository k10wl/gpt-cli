package initializers

import (
	"errors"
	"strings"
)

type Settings struct {
	Model  string
	System string
}

func loadSettings(config *Config, flags *Flags) (*Settings, error) {
	model, err := getModel(config, flags)
	if err != nil {
		return &Settings{}, err
	}

	system, err := getSystemMessage(config, flags)
	if err != nil {
		return &Settings{}, err
	}

	return &Settings{
		Model:  *model,
		System: system,
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

func getSystemMessage(config *Config, flags *Flags) (string, error) {
	system := flags.System

	if system == "" {
		return "", nil
	}

	message, exist := config.System[system]
	if !exist {
		errorMessage := "Requested system message (" + message + ") is not supported. Available: " + listAvailableSystem(config)
		return "", errors.New(errorMessage)
	}

	return message, nil
}

func listAvailableSystem(config *Config) string {
	available := make([]string, 0, len(config.System))
	for k := range config.System {
		available = append(available, k)
	}

	return strings.Join(available, ", ")
}
