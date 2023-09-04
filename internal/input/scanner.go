package input

import (
	"bufio"
	"fmt"
	"os"

	"cli/internal/initializers"
)

type Reader interface {
	Respond(message string) string
}

type AppState int

const (
	SystemRolePrefix    = "-> System"
	UserRolePrefix      = "-> User  "
	AssistantRolePrefix = "<- Assist"
)

func Scan(settings *initializers.Initialized, reader Reader) {
	scanner := bufio.NewScanner(os.Stdin)

	if settings.Flags.Session {
		clearConsole()

		if settings.Settings.Assistant != "" {
			fmt.Printf("%v\n", prefixMessage(SystemRolePrefix, settings.Settings.Assistant))
			fmt.Println(lineSeparator())
		}
	}

	for scanner.Scan() {
		message := scanner.Text()

		if !settings.Flags.Session {
			fmt.Println(reader.Respond(message))
			return
		}

		eraseConsoleLine()
		fmt.Printf("%v\n", prefixMessage(UserRolePrefix, message))
		fmt.Printf("%v\n", prefixMessage(AssistantRolePrefix, reader.Respond(message)))
		fmt.Println(lineSeparator())
	}
}
