package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/kimitoboku/tldr/tldr"
)

const (
	Version = "0.0.2"
)

var (
	platform string
	path     string
	docs     []string
)

type Config struct {
	Documetns []string
}

func fileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func tldrPrint(cmd string) {
	var cmdDoc string
	for _, doc := range docs {
		if fileExist(tldr.GenPath(doc, cmd)) {
			cmdDoc = tldr.PrintMd(tldr.GenPath(doc, cmd))
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
		fmt.Println(err.Error())
		content := []byte("{\n\"documetns\": [\"" + os.Getenv("HOME") + "/.config/tldr/tldr/pages/common\", \"" + os.Getenv("HOME") + "/.config/tldr/tldr/pages/linux\"]\n}")
		err = ioutil.WriteFile(os.Getenv("HOME")+"/.config/tldr/config.json", content, os.ModePerm)
		if err != nil {
			panic(err)
		}
		jsonFile, err = ioutil.ReadFile(os.Getenv("HOME") + "/.config/tldr/config.json")
		if err != nil {
			panic(err)
		}
	}
	err = json.Unmarshal(jsonFile, &data)
	if err != nil {
		fmt.Println(err.Error())
	}
	docs = data.Documetns
}

func main() {
	app := cli.NewApp()
	app.Name = "TL;DR"
	app.Usage = "tldr <command>"
	app.Version = Version
	app.Action = func(c *cli.Context) {
		tldrPrint(c.Args().First())
	}
	app.Run(os.Args)
}
