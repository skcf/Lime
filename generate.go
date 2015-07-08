package main

import (
	"github.com/codegangsta/cli"
	"io/ioutil"
	"os"
	"strings"
)

var commandG = cli.Command{
	Name:        "g",
	Usage:       "Generate a template file",
	Description: "",
	Action:      doG,
}

func doG(c *cli.Context) {
	switch {
	case len(c.Args()) == 1:
		generateTemplate(c, ".go", "Go/template.go")
		generateTemplate(c, ".py", "Python/template.py")
		generateTemplate(c, ".sh", "ShellScript/template.sh")
	case len(c.Args()) > 1:
		doHelp(c)
	default:
		doHelp(c)
	}
}

func generateTemplate(c *cli.Context, extension string, srcPath string) {
	home := os.Getenv("HOME") + "/.lime/Templates/"
	if strings.Contains(c.Args()[0], extension) {
		src, err := ioutil.ReadFile(home + srcPath)
		checkErr(err)
		err = ioutil.WriteFile("./"+c.Args()[0], src, 0644)
	}
}
