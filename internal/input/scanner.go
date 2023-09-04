package input

import (
	"bufio"
	"fmt"
	"os"
)

type Reader interface {
	Respond(message string) string
}

const (
	UserRolePrefix      = "-> User     "
	AssistantRolePrefix = "<- Assistant"
)

func Scan(reader Reader, session bool) {
	scanner := bufio.NewScanner(os.Stdin)

	if session {
		clearConsole()
	}

	for scanner.Scan() {
		// TODO this looks like shit, redo
		if session {
			eraseConsoleLine()
		}

		message := scanner.Text()

		if !session {
			fmt.Println(reader.Respond(message))
			return
		}

		fmt.Printf("%v\n", prefixMessage(UserRolePrefix, message))

		fmt.Printf("%v\n", prefixMessage(AssistantRolePrefix, reader.Respond(message)))

		fmt.Println(lineSeparator())
	}
}
