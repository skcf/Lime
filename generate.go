package main

import (
	"github.com/codegangsta/cli"
	"io/ioutil"
	"os"
	"regexp"
)

var commandG = cli.Command{
	Name:        "g",
	Usage:       "generate a template file",
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
	home := os.Getenv("HOME") + "/.lime/templates/generator-templates/"
	if checkExtension(c.Args()[0], extension) {
		src, err := ioutil.ReadFile(home + srcPath)
		checkErr(err)
		err = ioutil.WriteFile("./"+c.Args()[0], src, 0644)
	}
}

func checkExtension(filename string, extension string) bool {
	ex := "\\" + extension + "$"
	if m, _ := regexp.MatchString(ex, filename); !m {
		return false
	}
	return true
}
