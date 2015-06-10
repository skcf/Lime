package main

import (
  "os"
  "os/exec"
  "github.com/codegangsta/cli"
)

var Commands = []cli.Command{
    commandInit,
    commandList,
    commandSetup,
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

var commandSetup = cli.Command{
  Name: "setup",
  Usage: "",
  Description: "",
  Action: doSetup,
}

var commandHelp = cli.Command{
  Name: "help",
  Usage: "",
  Description: "",
  Action: doHelp,
}



func doInit(c *cli.Context) {
  home := os.Getenv("HOME")
  switch {
    case len(c.Args()) == 1:
      source := home + "/.lime/" + c.Args()[0]
      println("command ",c.Args()[0])
      cmd := exec.Command("cp","-r",source,".")
      _, err := cmd.Output()
      if err != nil {
        println(err.Error())
        return
      }
    case len(c.Args()) > 1:
      println("Error: Unknown command",c.Args()[0])
    default:
      doHelp(c)
  }
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
