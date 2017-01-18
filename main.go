package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/kimitoboku/eleking/tldr"
)

const (
	Version = "0.0.5"
)

var (
	platform   string
	path       string
	docs       []string
	existConfg bool
)

type Config struct {
	Documetns []string
}

func fileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func printCommandList() {
	for _, doc := range docs {
		fmt.Println(doc)
		files, err := ioutil.ReadDir(doc)
		if err != nil {
			fmt.Print(err.Error())
		}
		for _, file := range files {
			if !file.IsDir() {
				fmt.Println("  " + file.Name()[:len(file.Name())-3])
			}
		}
	}
}

func tldrPrint(cmd string) {
	var cmdDoc string
	for _, doc := range docs {
		if fileExist(tldr.GenPath(doc, cmd)) {
			cmdDoc = tldr.PrintMd(tldr.GenPath(doc, cmd))
			break
		}
	}
	if strings.Compare(cmd, "list") == 0 {
		printCommandList()
	} else if strings.Compare(cmdDoc, "") == 0 {
		tldrPrintRemote(cmd)
	} else {
		fmt.Print(cmdDoc)
	}
}

func tldrPrintRemote(cmd string) {
	var cmdDoc string
	plist := []string{"common", "linux"}
	for _, pl := range plist {
		cmdDoc = tldr.HttpGetMd(pl, cmd)
		if strings.Compare(cmdDoc, "") != 0 {
			break
		}
	}

	if strings.Compare(cmdDoc, "") == 0 {
		fmt.Println("Command Not Found")
	} else {
		fmt.Print(cmdDoc)
	}
}

func init() {
	var data Config
	jsonFile, err := ioutil.ReadFile(os.Getenv("HOME") + "/.config/tldr/config.json")
	if err != nil {
		existConfg = false
	} else {
		existConfg = true
		err = json.Unmarshal(jsonFile, &data)
		if err != nil {
			fmt.Println(err.Error())
		}
		docs = data.Documetns
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "eleking"
	app.Usage = "eleking <command>"
	app.Version = Version
	app.Action = func(c *cli.Context) {
		if existConfg {
			tldrPrint(c.Args().First())
		} else {
			tldrPrintRemote(c.Args().First())
		}
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err.Error())
	}
}
