package main

import (
	"os"

	"github.com/urfave/cli"
  "github.com/coolopsio/coolops/info"
)

func main() {

	app := cli.NewApp()
	app.Name = info.Name
	app.Version = info.Version
	app.Author = "coolopsio"
	app.Email = ""
	app.Usage = ""

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}
