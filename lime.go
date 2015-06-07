package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Lime"
	app.Version = Version
	app.Usage = ""
	app.Author = "Souichi"
	app.Email = "sk.cf.msc@gmail.com"
	app.Action = doMain
	app.Commands = Commands
	app.Run(os.Args)
}

func doMain(c *cli.Context) {
	println ("hello world")
}
