package input

import gpt_client "github.com/k10wl/gpt-client"

var messagesCache = []gpt_client.Message{}

func ApplySystemMessage(message string) {
	messagesCache = append([]gpt_client.Message{
		{Role: "system", Content: message},
	},
		messagesCache...,
	)
}
