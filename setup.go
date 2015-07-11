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
	generatorTemplate := os.Getenv("GOPATH") + "/src/github.com/skcf/Lime/templates/generator-templates"
	initTemplate := os.Getenv("GOPATH") + "/src/github.com/skcf/Lime/templates/init-templates"
	cmd1 := exec.Command("cp", "-r", generatorTemplate, limePath)
	_, err := cmd1.Output()
	if err != nil {
		println(err.Error())
	}
	cmd2 := exec.Command("cp", "-r", initTemplate, limePath)
	_, err = cmd2.Output()
	if err != nil {
		println(err.Error())
	}
}
