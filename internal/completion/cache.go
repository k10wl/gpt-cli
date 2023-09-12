package completion

import (
	"encoding/json"
	"os"

	gpt_client "github.com/k10wl/gpt-client"
)

func (c *Client) updateSession(messages *[]gpt_client.Message) error {
	file, err := os.OpenFile("session.cache", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(messages)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) readSession() (*[]gpt_client.Message, error) {
	file, err := os.Open("session.cache")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var messages []gpt_client.Message
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&messages); err != nil {
		return nil, err
	}
	return &messages, nil
}
