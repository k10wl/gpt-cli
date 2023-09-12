package completion

import (
	"log"
	"os"

	gpt_client "github.com/k10wl/gpt-client"
)

type Client struct {
	client        *gpt_client.Client
	systemMessage *gpt_client.Message
}

var includedMessages = make([]gpt_client.Message, 0)

func New() *Client {
	return &Client{
		client: gpt_client.NewClient(os.Getenv("OPENAI_API_KEY")),
	}
}

func (c *Client) Complete(prompt string) (string, error) {
	messages, _ := c.readSession()
	userMessage := gpt_client.Message{Role: "user", Content: prompt}
	gptMessages := append(*messages, userMessage)
	history, err := c.client.BuildHistory(&gptMessages)
	if err != nil {
		return "", err
	}
	res, err := c.client.TextCompletion(history)
	if err != nil {
		log.Fatalf("Error from GPT client: %+v\n", err)
	}
	updatedHistory := append(*history, res.Choices[0].Message)
	c.updateSession(&updatedHistory)
	return res.Choices[0].Message.Content, err
}

func (c *Client) SystemMessage(prompt string) {
	c.systemMessage = &gpt_client.Message{
		Role:    "system",
		Content: prompt,
	}
}

func (c *Client) DropHistory() {
	var clearCache []gpt_client.Message
	if c.systemMessage != nil {
		clearCache = make([]gpt_client.Message, 1)
		clearCache[0] = *c.systemMessage
	} else {
		clearCache = make([]gpt_client.Message, 0)
	}
	c.updateSession(&clearCache)
}
