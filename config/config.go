package config

import (
	"os"
)

type Global struct {
	Flags
	AssistantMessage string
}

func New() *Global {
	flags := readFlags()
	loadEnv(flags.Path)

	// XXX should this be here?
	if flags.ListAssistants {
		printAssistants(flags.Path)
		os.Exit(0)
	}

	return &Global{
		Flags:            *flags,
		AssistantMessage: getAssistantMessage(flags.Path, flags.Assistant),
	}
}
