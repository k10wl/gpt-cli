package input

var systemMessages = map[string]string{
	"nvim": "Help with Neovim and VIM. Explain commands and answer questions by providing clear instructions",
}

func RetreiveSystemMessage(text string) string {
	instruction, ok := systemMessages[text]

	if !ok {
		panic("System message does not exits")
	}

	return instruction
}
