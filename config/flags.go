package config

import (
	"flag"
	"log"
	"os"
)

type Flags struct {
	Path           string
	Assistant      string
	Session        bool
	Restore        bool
	ListAssistants bool
}

func readFlags() *Flags {
	currentPath, err := os.Getwd()
	if err != nil {
		currentPath = ""
	}

	var path string
	flag.StringVar(&path, "path", currentPath, "path to execution binary")
	if path == "" {
		log.Fatal("Cannot read current path")
	}

	var assistant string
	flag.StringVar(&assistant, "assistant", "", "specify predefined assistant type")
	flag.StringVar(&assistant, "a", "", "specify predefined assistant type (shorthand)")

	var session bool
	flag.BoolVar(&session, "session", false, "start chat session")
	flag.BoolVar(&session, "s", false, "start chat session (shorthand)")

	var restore bool
	flag.BoolVar(&restore, "restore", false, "restore last session if available")
	flag.BoolVar(&restore, "r", false, "restore last session if available (shorthand)")

	var listAssistants bool
	flag.BoolVar(&listAssistants, "list-assistants", false, "list available assistants")
	flag.BoolVar(&listAssistants, "la", false, "list available assistants (shorthand)")

	flag.Parse()

	return &Flags{
		Path:           path,
		Assistant:      assistant,
		Session:        session,
		Restore:        restore,
		ListAssistants: listAssistants,
	}
}
