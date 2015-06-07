package main

import (
  "os"
  "os/exec"
  "github.com/codegangsta/cli"
)

var Commands = []cli.Command{
    commandInit,
    commandList,
    commandHelp,
}

var commandInit = cli.Command{
  Name: "init",
  Usage: "",
  Description: "",
  Action: doInit,
}

var commandList = cli.Command{
  Name: "list",
  Usage: "",
  Description: "",
  Action: doList,
}

var commandHelp = cli.Command{
  Name: "help",
  Usage: "",
  Description: "",
  Action: doHelp,
}



func doInit(c *cli.Context) {
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

func doList(c *cli.Context) {
  home := os.Getenv("HOME")
  targetPath := home + "/.lime"

  cmd := exec.Command("ls",targetPath)

  stdout, err := cmd.Output()
  if err != nil {
    println(err.Error())
    return
  }
  println(string(stdout))
}
