package main

import (
	"github.com/guumaster/go-tts/cmd"
)

// Inspired by https://github.com/hegedustibor/htgo-tts
// For excusator:
// text = "Cuando averigüemos por qué se cae el proceso del sistema de colas del acelerador de transacciones"

var version = "dev"

func main() {
	cmd.RunCLI(version)
}
