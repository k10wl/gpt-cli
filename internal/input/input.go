package input

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"cli/internal/completion"
)

const (
	linebreaksToDropHistory int8 = 3
	idle                         = iota
	processing
)

type InputScanner struct {
	currentState             int
	messagedDuringProcessing bool
	client                   *completion.Client
	scanner                  *bufio.Scanner
	emptyLines               int8
	assistantName            string
	userName                 string
	processingMessage        string
}

func Create(client *completion.Client, assistantName string) *InputScanner {
	if assistantName == "" {
		assistantName = "default"
	}
	userName, assistantName := addSpaces("user", assistantName)
	processingMessage := prefixMessage(assistantName, "Processing, please wait...", GreenColor)

	return &InputScanner{
		currentState:             idle,
		messagedDuringProcessing: false,
		client:                   client,
		scanner:                  bufio.NewScanner(os.Stdin),
		emptyLines:               0,
		assistantName:            assistantName,
		userName:                 userName,
		processingMessage:        processingMessage,
	}
}

func (i *InputScanner) Session(restored bool) {
	i.prepareSessionScreen(restored)
	for i.scanner.Scan() {
		go i.process()
	}
}

func (i *InputScanner) SinglePrompt() {
	i.emptyLines = linebreaksToDropHistory
	i.scanner.Scan()
	completion, err := i.client.Complete(i.scanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(completion)
}

func (i *InputScanner) process() {
	text := i.scanner.Text()
	if i.messagedDuringProcessing {
		eraseTextLines(i.processingMessage)
	}
	if i.currentState == processing {
		i.messagedDuringProcessing = true
		eraseTextLines(text)
		fmt.Println(i.processingMessage)
		return
	}

	if text == "" {
		i.handleEmptyText()
		return
	}

	// TODO prefix should always be on screen
	formatUserMessage(i.userName, text)
	i.emptyLines = 0
	i.currentState = processing
	i.messagedDuringProcessing = false
	completion, err := i.client.Complete(text)
	if err != nil {
		log.Fatal(err)
	}
	if i.messagedDuringProcessing {
		eraseTextLines(i.processingMessage)
	}
	fmt.Println(prefixMessage(i.assistantName, completion, GreenColor))
	fmt.Println(lineSeparator())
	i.messagedDuringProcessing = false
	i.currentState = idle
}

func (i *InputScanner) handleEmptyText() {
	if i.emptyLines == linebreaksToDropHistory {
		i.client.DropHistory()
		i.emptyLines = 0
		i.prepareSessionScreen(false)
		return
	}

	if i.emptyLines > 0 {
		eraseLinesCount(2)
	} else {
		eraseLinesCount(1)
	}

	message := "Hit enter " +
		strconv.FormatInt(int64(linebreaksToDropHistory-i.emptyLines), 10) +
		" more time(s) to drop history"

	fmt.Println(prefixMessage(i.assistantName, message, GreenColor))

	i.emptyLines += 1
}

func (i *InputScanner) prepareSessionScreen(restored bool) {
	sufix := " assistant"
	if restored {
		sufix += " (restored)"
	}
	clearTerminal()
	fmt.Println(centerMessage(i.assistantName+sufix) + "\n")
}
