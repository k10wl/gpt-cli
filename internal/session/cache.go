package session

import gpt_client "github.com/k10wl/gpt-client"

var messagesCache = []gpt_client.Message{}

func getCache() *[]gpt_client.Message {
	return &messagesCache
}

func updateCache(updated *[]gpt_client.Message) {
	messagesCache = *updated
}

func cacheMessage(message gpt_client.Message) []gpt_client.Message {
	messagesCache = append(messagesCache, message)
	return messagesCache
}
