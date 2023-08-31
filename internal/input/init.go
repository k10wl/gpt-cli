package input

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	gpt_client "github.com/k10wl/gpt-client"
)

func Start() {
	scanner := bufio.NewScanner(os.Stdout)

	apiKey := os.Getenv("OPENAI_API_KEY")
	client := gpt_client.NewClient(apiKey)

	shouldReturn := handleFlags(client)

	if shouldReturn {
		return
	}

	for scanner.Scan() {
		message := scanner.Text()
		submitTextCompletion(client, message)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func handleFlags(client *gpt_client.Client) bool {
	shouldBreak := false

	config := flag.String("c", "", "system message config")
	message := flag.String("m", "", "input message")
	flag.Parse()

	if *config != "" {
		ApplySystemMessage(RetreiveSystemMessage(*config))
	}

	if *message != "" {
		shouldBreak = true
		submitTextCompletion(client, *message)
	}

	return shouldBreak
}

func submitTextCompletion(client *gpt_client.Client, message string) {
	messagesCache = append(messagesCache, gpt_client.Message{Role: "user", Content: message})

	history, err := client.BuildHistory(&messagesCache)
	if err != nil {
		panic(err)
	}

	messagesCache = *history

	res, err := client.TextCompletion(&messagesCache)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Choices[0].Message.Content)
}
