package main

import (
    "os"
    "os/exec"
    "fmt"
    "log"
    "io/ioutil"
    "regexp"
    "github.com/codegangsta/cli"
)

var Commands = []cli.Command{
    commandInit,
    commandLs,
    commandSetup,
    commandHelp,
}

var commandInit = cli.Command{
    Name: "init",
    Usage: "",
    Description: "",
    Action: doInit,
}

var commandLs = cli.Command{
    Name: "ls",
    Usage: "",
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

func doLs(c *cli.Context) {
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
