package main

import (
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
)

var commandSetup = cli.Command{
	Name:        "setup",
	Usage:       "setup template directory",
	Description: "",
	Action:      doSetup,
}

func doSetup(c *cli.Context) {
	home := os.Getenv("HOME")
	if exist(home + "/.lime") {
		copyTemplates()
	} else {
		targetPath := home + "/.lime"
		cmd := exec.Command("mkdir", targetPath)
		stdout, err := cmd.Output()
		if err != nil {
			println(err.Error())
			return
		}
		println(string(stdout))
		copyTemplates()
	}
}

func copyTemplates() {
	home := os.Getenv("HOME")
	limePath := home + "/.lime/"
	templatePath := os.Getenv("GOPATH") + "/src/github.com/skcf/Lime/Templates"
	cmd := exec.Command("cp", "-r", templatePath, limePath)
	_, err := cmd.Output()
	if err != nil {
		println(err.Error())
	}
}
