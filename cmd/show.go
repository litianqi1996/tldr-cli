package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"strings"
)

// todo : bug in tldr   head and tail

var (
	titleShow       = color.New(color.FgWhite, color.Bold).PrintlnFunc()
	referShow       = color.New(color.FgWhite).PrintlnFunc()
	instructionShow = color.New(color.FgGreen).PrintlnFunc()
	commandShow     = color.New(color.FgCyan).PrintFunc()
	argsShow        = color.New(color.FgRed).PrintFunc()
)

func ShowItem(command string) {
	items := strings.Split(strings.Trim(command, "`"), " ")
	for _, item := range items {
		if strings.Contains(item, "{{") {
			commandShow(" " + strings.Trim(item, "{{}}") + " ") //todo: bug in head and tail
		} else {
			argsShow(" " + item)
		}
	}
}

func Show(pages []byte) {

	lhs := strings.Split(string(pages), "\n")
	for _, value := range lhs {
		switch {
		case strings.HasPrefix(value, "-"):
			instructionShow(value)
		case strings.HasPrefix(value, "`"):
			ShowItem(value)
			fmt.Println("\n")
		case strings.HasPrefix(value, "#"):
			titleShow(strings.Split(value, "#")[1] + "\n")
		case strings.HasPrefix(value, ">"):
			referShow(strings.Split(value, ">")[1] + "\n")
		}
	}
}

func Getpage(cmd string) error {
	cmdpath, err := SearchPages(cmd, RepoPath, Language)
	if err != nil {
		return err
	}
	page, err := ioutil.ReadFile(cmdpath)
	if err != nil {
		return err
	}
	Show(page)
	return nil
}
