package main

import (
	"log"
	"os"

	"github.com/ichtrojan/chow/chow"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:    "chow",
		Usage:   "Parse the JSON file",
		Action:  chow.HandleFunc,
		Version: "1.2.0",
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
