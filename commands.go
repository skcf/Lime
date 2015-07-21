package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"os"
)

var Commands = []cli.Command{
	commandInit,
	commandG,
	commandLs,
	commandSetup,
	commandHelp,
}

var commandLs = cli.Command{
	Name:        "ls",
	Usage:       "list lime templates",
	Description: "",
	Action:      doLs,
}

var commandHelp = cli.Command{
	Name:        "help",
	Usage:       "show help",
	Description: "",
	Action:      doHelp,
}

func doLs(c *cli.Context) {
	switch {
	case len(c.Args()) == 0:
		home := os.Getenv("HOME")
		generatorTemplatesPath := home + "/.lime/generator-templates"
		initTemplatesPath := home + "/.lime/init-templates"
		initScriptsPath := home + "/.lime/init-scripts"

		var generatorTemplateList []string
		var initTemplateList []string
		var scriptList []string

		generatorTemplates, _ := ioutil.ReadDir(generatorTemplatesPath)
		for _, f := range generatorTemplates {
			if f.IsDir() == true {
				generatorTemplateList = append(generatorTemplateList, f.Name())
			}
		}

		initTemplates, _ := ioutil.ReadDir(initTemplatesPath)
		for _, f := range initTemplates {
			if f.IsDir() == true {
				initTemplateList = append(initTemplateList, f.Name())
			}
		}

		initScripts, _ := ioutil.ReadDir(initScriptsPath)
		for _, f := range initScripts {
			scriptList = append(scriptList, f.Name())
		}

		fmt.Println("\n*** Generator Templates ***")
		for _, i := range generatorTemplateList {
			fmt.Println(i)
		}

		fmt.Println("\n*** Init Templates ***")
		for _, i := range initTemplateList {
			fmt.Println(i)
		}

		fmt.Println("\n*** Init Scripts ***")
		for _, i := range scriptList {
			fmt.Println(i)
		}
		fmt.Println("")
	case len(c.Args()) == 1:
		home := os.Getenv("HOME")
		targetPath := home + "/.lime/" + c.Args()[0]

		fmt.Println("\n*** Lime Template " + c.Args()[0] + " ***")

		files, _ := ioutil.ReadDir(targetPath)
		for _, f := range files {
			fmt.Println(f.Name())
		}
	}
}
