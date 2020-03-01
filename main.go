package main

import (
	"github.com/guumaster/go-tts/cmd"
)

// Inspired by https://github.com/hegedustibor/htgo-tts

var version = "dev"

func main() {
	cmd.RunCLI(version)
}
