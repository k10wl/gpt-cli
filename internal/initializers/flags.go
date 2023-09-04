package initializers

import (
	"flag"
)

type Flags struct {
	Path    string
	System  string
	Session bool
	Model   string
}

func readFlags() *Flags {
	path := flag.String("path", "", "app root path")
	system := flag.String("system", "", "system message")
	session := flag.Bool("session", false, "start session")
	model := flag.String("model", "", "language model")

	flag.Parse()

	return &Flags{
		Path:    *path,
		System:  *system,
		Session: *session,
		Model:   *model,
	}
}
