package main

import (
    "os"
    "os/exec"
    "fmt"
    "log"
    "io/ioutil"
    "regexp"
    "strings"
    "github.com/codegangsta/cli"
)

var Commands = []cli.Command{
    commandInit,
    commandG,
    commandLs,
    commandSetup,
    commandHelp,
}

var commandInit = cli.Command{
    Name: "init",
    Usage: "copy template directory and run \"lime.sh\" in template directory",
    Description: "",
    Action: doInit,
}

var commandG = cli.Command{
    Name: "g",
    Usage: "Generate a template file",
    Description: "",
    Action: doG,
}

var commandLs = cli.Command{
    Name: "ls",
    Usage: "list lime templates",
    Description: "",
    Action: doLs,
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
            if shell(c.Args()[0]) == true {
                copy744(source,"./" + c.Args()[0])
                out, err := exec.Command("./" + c.Args()[0]).Output()
                if err != nil {
                    log.Fatal(err)
                }
                fmt.Printf("%s",out)
                removeFile(c.Args()[0])
            } else {
                files, _ := ioutil.ReadDir(source)
                for _, f := range files {
                    if f.Name() == "lime.sh" {
                        copy744(source + "/" + f.Name(),"./lime.sh")
                    } else {
                        copyFile(source + "/"+ f.Name())
                    }
                }
                if exist("lime.sh") == true {
                    out, err := exec.Command("./lime.sh").Output()
                    if err != nil {
                        log.Fatal(err)
                    }
                    fmt.Printf("%s",out)
                    removeFile("lime.sh")
                }
            }
        case len(c.Args()) > 1:
            println("Error: Unknown command",c.Args()[0])
        default:
            doHelp(c)
    }
}

func shell(name string) bool {
    if m, _ := regexp.MatchString(".sh$", name); !m {
        return false
    }
    return true
}

func copyFile(src string){
    cmd := exec.Command("cp","-r",src,".")
    _, err := cmd.Output()
    if err != nil {
        println(err.Error())
        return
    }
}

func copy744(src string, dst string) {
	data, err := ioutil.ReadFile(src)
	checkErr(err)
	err = ioutil.WriteFile(dst,data,0744)
	checkErr(err)
}

func checkErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func exist(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil
}

func removeFile(file string) {
    cmd := exec.Command("rm","-f","./lime.sh")
    _, err := cmd.Output()
    if err != nil {
        println(err.Error())
        return
    }
}

func doG(c *cli.Context) {
    generateTemplate(c,".go","Go/template.go")
    generateTemplate(c,".py","Python/template.py")
}

func generateTemplate(c *cli.Context,extension string, srcPath string) {
  templatePath := os.Getenv("GOPATH") + "/src/github.com/skcf/Lime/Templates/"
  if strings.Contains(c.Args()[0], extension) == true {
      src, err:= ioutil.ReadFile( templatePath + srcPath )
      checkErr(err)
      err = ioutil.WriteFile("./" + c.Args()[0],src,0644)
  }
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
                  templates = append(templates,f.Name())
                } else {
                  scripts = append(scripts,f.Name())
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
