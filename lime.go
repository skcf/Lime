package main

import (
	"os"

	"github.com/codegangsta/cli"
)

var limeHelpTemplate = `
Lime is a tool for generating template files and building environment.

USAGE:
   {{.Name}} {{if .Flags}}[global options] {{end}}command{{if .Flags}} [command options]{{end}} [arguments...]

VERSION:
   {{.Version}}{{if len .Authors}}

AUTHOR(S):
   {{range .Authors}}{{ . }}{{end}}{{end}}

COMMANDS:
   {{range .Commands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
   {{end}}{{if .Flags}}

GLOBAL OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{end}}
`

func main() {
	initTemplate()
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
	switch {
	case len(c.Args()) == 1:
		println("command ",c.Args()[0])
	case len(c.Args()) > 1:
		println("Error: Unknown command",c.Args()[0])
	default:
		doHelp(c)
	}
}

func doHelp(c *cli.Context) {
	cli.AppHelpTemplate = limeHelpTemplate
	cli.ShowAppHelp(c)
}


func initTemplate() {
	cli.AppHelpTemplate = limeHelpTemplate
}
