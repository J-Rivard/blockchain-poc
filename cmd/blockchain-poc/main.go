package main

import (
	"os"

	"github.com/J-Rivard/blockchain-poc/internal/logging"
	"github.com/rs/zerolog"
)

func main() {
	logger, err := logging.New(zerolog.ConsoleWriter{Out: os.Stderr}, logging.Debug)
	if err != nil {
		panic(err)
	}
	logger.LogDebug(map[string]string{"test": "test"})
}
