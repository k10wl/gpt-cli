package initializers

import (
	"fmt"
	"os"

	gpt_client "github.com/k10wl/gpt-client"
)

type Initialized struct {
	Flags     *Flags
	Settings  *Settings
	GPTClient *gpt_client.Client
}

func LoadAll() *Initialized {
	flags := readFlags()

	err := LoadEnv(flags)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	config, err := readConfigFile(flags)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	settings, err := loadSettings(config, flags)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	client := gpt_client.NewClient(apiKey)

	return &Initialized{
		Flags:     flags,
		Settings:  settings,
		GPTClient: client,
	}
}
