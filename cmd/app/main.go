package main

import (
	"cli/internal/initializers"
	"cli/internal/input"
	"cli/internal/session"
)

func main() {
	data := initializers.LoadAll()

	sessionReader := session.NewSession(data)
	input.Scan(data, sessionReader)
}
