package main

import (
	"os"
  "os/exec"
	"github.com/codegangsta/cli"
)

var commandSetup = cli.Command{
    Name: "setup",
    Usage: "",
    Description: "",
    Action: doSetup,
}

func doSetup(c *cli.Context) {
    home := os.Getenv("HOME")
    targetPath := home + "/.lime"
    cmd := exec.Command("mkdir",targetPath)
    stdout, err := cmd.Output()
    if err != nil {
        println(err.Error())
        return
    }
    println(string(stdout))
}
