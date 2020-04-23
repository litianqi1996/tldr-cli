package cmd

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fatih/color"
)

var (
	titleShow       = color.New(color.FgWhite, color.Bold).PrintlnFunc()
	referShow       = color.New(color.FgWhite).PrintlnFunc()
	instructionShow = color.New(color.FgGreen).PrintlnFunc()
	commandShow     = color.New(color.FgCyan).PrintFunc() //not println
	argsShow        = color.New(color.FgRed).PrintFunc()
)

func Getpage(cmd string) error {
	cmdpath, err := SearchPages(cmd, RepoPath, Language)
	if err != nil {
		return err
	}
	page, err := ioutil.ReadFile(cmdpath)
	if err != nil {
		return err
	}
	Show(string(page))
	return nil
}

func Show(pages string) {

	fs := format(pages)

	for _, value := range fs {
		switch {
		case strings.HasPrefix(value, "-"):
			instructionShow(value)
		case strings.HasPrefix(value, "`"):
			showItem(value)
		case strings.HasPrefix(value, "#"):
			titleShow("\n" + strings.TrimLeft(value, "#"))
		case strings.HasPrefix(value, ">"):
			referShow(strings.TrimLeft(value, ">"))
		case value == "":
			fmt.Print("\n")
		}
	}
}

func showItem(command string) {
	items := strings.Split(strings.Trim(command, "`"), " ")
	for _, item := range items {
		if strings.Contains(item, "{{") {
			commandShow(" " + strings.ReplaceAll(strings.ReplaceAll(item, "{{", ""), "}}", "") + " ")
		} else {
			argsShow(" " + item)
		}
	}
	fmt.Print("\n")
}

func format(page string) []string {
	lhs := strings.Split(string(page), "\n")
	for i, value := range lhs {
		if strings.HasPrefix(value, "-") {
			if lhs[i+1] == "" {
				lhs = append(lhs[:i+1], lhs[i+2:]...)
			}
		}
	}
	return lhs
}
