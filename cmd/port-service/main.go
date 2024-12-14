package main

import (
	"log"
	"os"

	"github.com/platinumscatter/port-service/internal/config"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run() error {
	_ = config.Read()
	return nil
}
