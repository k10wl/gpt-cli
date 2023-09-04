package main

import (
	"cli/internal/initializers"
	"cli/internal/input"
	"cli/internal/session"
)

func main() {
	data := initializers.LoadAll()

	s := session.NewSession(data)
	input.Scan(s, data.Flags.Session)
}
