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
	Usage:       "",
	Description: "",
	Action:      doHelp,
}

func doLs(c *cli.Context) {
	switch {
	case len(c.Args()) == 0:
		home := os.Getenv("HOME")
		targetPath := home + "/.lime"

		var templates []string
		var scripts []string

		files, _ := ioutil.ReadDir(targetPath)
		for _, f := range files {
			if f.IsDir() == true {
				templates = append(templates, f.Name())
			} else {
				scripts = append(scripts, f.Name())
			}
		}

		fmt.Println("\n*** Lime Template Directories ***")
		for _, i := range templates {
			fmt.Println(i)
		}

		fmt.Println("\n*** Lime Template Scripts ***")
		for _, i := range scripts {
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
