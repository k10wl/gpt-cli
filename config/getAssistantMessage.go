package config

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"strings"
)

const assistantsFile = "assistants.json"

func getAssistantMessage(dir string, name string) string {
	if name == "" {
		return ""
	}

	assistants, err := assistantsReader(dir)
	if err != nil {
		log.Printf("Skipping system message due to error upon reading available assistants: %+v", err)
		return ""
	}

	assistantMessage, exists := assistants[name]
	if !exists {
		log.Fatal("unknown assistant `" + name + "`, available assistants: " + strings.Join((*listAvailableAssistants(assistants)), "; "))
	}

	return assistantMessage
}

func assistantsReader(dir string) (map[string]string, error) {
	var assistants map[string]string

	assistantsPath := path.Join(dir, assistantsFile)
	b, error := os.ReadFile(assistantsPath)
	if error != nil {
		return nil, error
	}

	if err := json.Unmarshal(b, &assistants); err != nil {
		return nil, err
	}

	return assistants, nil
}

func listAvailableAssistants(assistants map[string]string) *[]string {
	list := make([]string, len(assistants))

	i := 0
	for key := range assistants {
		list[i] = key
		i = i + 1
	}

	return &list
}