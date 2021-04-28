package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/simonsargeant/telemeterm/internal/command"
)

func main() {
	root := command.NewRoot()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
