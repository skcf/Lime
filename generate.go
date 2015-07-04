package main

import (
    "os"
    "io/ioutil"
    "strings"
    "github.com/codegangsta/cli"
)

var commandG = cli.Command{
    Name: "g",
    Usage: "Generate a template file",
    Description: "",
    Action: doG,
}

func doG(c *cli.Context) {
    switch {
    case len(c.Args()) == 1:
        generateTemplate(c,".go","Go/template.go")
        generateTemplate(c,".py","Python/template.py")
    case len(c.Args()) > 1:
        doHelp(c)
    default:
        doHelp(c)
    }
}

func generateTemplate(c *cli.Context,extension string, srcPath string) {
  templatePath := os.Getenv("GOPATH") + "/src/github.com/skcf/Lime/Templates/"
  if strings.Contains(c.Args()[0], extension) == true {
      src, err:= ioutil.ReadFile( templatePath + srcPath )
      checkErr(err)
      err = ioutil.WriteFile("./" + c.Args()[0],src,0644)
  }
}
