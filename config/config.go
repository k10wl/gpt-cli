package config

type Global struct {
	Flags
	AssistantMessage string
}

func New() *Global {
	flags := readFlags()
	loadEnv(flags.Path)

	return &Global{
		Flags:            *flags,
		AssistantMessage: getAssistantMessage(flags.Path, flags.Assistant),
	}
}
