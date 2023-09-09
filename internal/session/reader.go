package session

import (
	"log"

	"cli/internal/initializers"

	gpt_client "github.com/k10wl/gpt-client"
)

type Reader interface {
	Respond(text string) string
}

type Current struct {
	initializers.Initialized
}

func NewSession(data *initializers.Initialized) *Current {
	if data.Settings.Assistant != "" {
		applyAssistantMessage(data.Settings.Assistant)
	}

	return &Current{
		*data,
	}
}

func applyAssistantMessage(message string) {
	cacheMessage(gpt_client.Message{
		Role:    "system",
		Content: message,
	})
}

func (c *Current) Respond(text string) string {
	cacheMessage(gpt_client.Message{Role: "user", Content: text})

	history, err := c.GPTClient.BuildHistory(getCache())
	if err != nil {
		log.Fatal(err)
	}

	res, err := c.GPTClient.TextCompletion(history)
	if err != nil {
		log.Fatal(err)
	}

	cacheMessage(res.Choices[0].Message)

	return res.Choices[0].Message.Content
}
