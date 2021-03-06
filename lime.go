package main

import (
	"github.com/codegangsta/cli"
	"os"
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
	// app.Action = doMain
	app.Commands = Commands
	app.Run(os.Args)
}

func doHelp(c *cli.Context) {
	cli.AppHelpTemplate = limeHelpTemplate
	cli.ShowAppHelp(c)
}

func initTemplate() {
	cli.AppHelpTemplate = limeHelpTemplate
}
