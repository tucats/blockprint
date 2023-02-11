package main

import (
	"fmt"
	"os"

	"github.com/tucats/gopackages/app-cli/app"
)

func main() {

	app.MakePrivate("logon")
	app.MakePrivate("insecure")
	app.MakePrivate("log")
	app.MakePrivate("log-file")
	app.MakePrivate("quiet")
	app.MakePrivate("format")

	app := app.New("blockprint: Print large text using github.com/common-nighthawk/go-figure").SetVersion(1, 0, 2)

	if err := app.Run(grammar, os.Args); err != nil {
		fmt.Println("Error,", err)
	}
}
