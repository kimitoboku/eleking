package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kimitoboku/tldr/tldr"
)

const (
	Version = "0.0.1"
)

var (
	platform string
	path     string
)

func main() {
	app := cli.NewApp()
	app.Name = "TL;DR"
	app.Usage = "tldr <command>"
	app.Version = Version
	app.Action = func(c *cli.Context) {
		path += os.Getenv("HOME") + "/.config/tldr"
		fmt.Print(tldr.PrintMd(tldr.GenPath(path, platform, c.Args().First())))
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "platform,p",
			Value:       "common",
			Usage:       "Search Platform",
			Destination: &platform,
		},
	}
	app.Run(os.Args)
}
