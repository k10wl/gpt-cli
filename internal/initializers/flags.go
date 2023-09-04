package initializers

import (
	"flag"
)

type Flags struct {
	Path      string
	Assistant string
	Session   bool
	Model     string
}

func readFlags() *Flags {
	var path string
	flag.StringVar(&path, "path", "", "app root path")
	flag.StringVar(&path, "p", "", "app root path (shorthand)")

	var assistant string
	flag.StringVar(&assistant, "assistant", "", "assistant message")
	flag.StringVar(&assistant, "a", "", "assistant message (shorthand)")

	var session bool
	flag.BoolVar(&session, "session", false, "start session")
	flag.BoolVar(&session, "s", false, "start session (shorthand)")

	var model string
	flag.StringVar(&model, "model", "", "language model")
	flag.StringVar(&model, "m", "", "language model (shorthand)")

	flag.Parse()

	return &Flags{
		Path:      path,
		Assistant: assistant,
		Session:   session,
		Model:     model,
	}
}
