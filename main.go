package main

import (
	"github.com/fatih/color"
	"github.com/litianqi1996/tldr/cmd"
	"github.com/urfave/cli/v2"
	"os"
)

const VERSION = "0.1"

func main() {
	app := cli.NewApp()
	app.Version = VERSION
	app.Name = "TLDR"
	app.Usage = "Too Long; Didn't Read"
	app.HideVersion = true
	app.HideHelpCommand = true
	app.HideHelp = true
	app.UsageText = "tldr [command]"

	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:     "update",
			Aliases:  []string{"u"},
			Usage:    "update tldr pages from gitrepo",
			Required: false,
		},
	}

	//  do not return err for https://github.com/urfave/cli/issues/707
	app.Before = func(c *cli.Context) error {
		err := cmd.StartUp()
		if err != nil {
			ErrShow(err)
			_ = cli.ShowAppHelp(c)
		}
		return nil
	}

	app.Action = func(c *cli.Context) error {
		if c.Bool("update") {
			err := cmd.UpdateRepo()
			if err != nil {
				return err
			}
		}
		if c.NArg() > 0 {
			err := cmd.Getpage(c.Args().Get(0))
			if err != nil {
				return err
			}
		} else {
			_ = cli.ShowAppHelp(c)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		ErrShow(err)
	}
}

var ErrShow = color.New(color.FgRed).PrintlnFunc()
