package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
)

var commandInit = cli.Command{
	Name:        "init",
	Usage:       "copy template directory and run \"lime.sh\" in template directory",
	Description: "",
	Action:      doInit,
}

func doInit(c *cli.Context) {
	home := os.Getenv("HOME")
	switch {
	case len(c.Args()) == 1:
		source := home + "/.lime/init-templates/" + c.Args()[0]
		if shell(c.Args()[0]) {
			copy744(source, "./"+c.Args()[0])
			out, err := exec.Command("./" + c.Args()[0]).Output()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s", out)
			removeFile(c.Args()[0])
		} else {
			files, _ := ioutil.ReadDir(source)
			for _, f := range files {
				if f.Name() == "lime.sh" {
					copy744(source+"/"+f.Name(), "./lime.sh")
				} else {
					copyFile(source + "/" + f.Name())
				}
			}
			if exist("lime.sh") {
				out, err := exec.Command("./lime.sh").Output()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%s", out)
				removeFile("lime.sh")
			}
		}
	case len(c.Args()) > 1:
		println("Error: Unknown command", c.Args()[0])
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

func copyFile(src string) {
	cmd := exec.Command("cp", "-r", src, ".")
	_, err := cmd.Output()
	if err != nil {
		println(err.Error())
		return
	}
}

func copy744(src string, dst string) {
	data, err := ioutil.ReadFile(src)
	checkErr(err)
	err = ioutil.WriteFile(dst, data, 0744)
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
	cmd := exec.Command("rm", "-f", "./lime.sh")
	_, err := cmd.Output()
	if err != nil {
		println(err.Error())
		return
	}
}
