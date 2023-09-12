package main

import (
	"cli/config"
	"cli/internal/completion"
	"cli/internal/input"
)

func main() {
	config := config.New()
	client := completion.New()
	input := input.Create(client, config.Assistant)

	if config.AssistantMessage != "" {
		client.SystemMessage(config.AssistantMessage)
	}

	if !config.Restore {
		client.DropHistory()
	}

	if config.Session {
		input.Session(config.Restore)
		return
	}

	input.SinglePrompt()
}
