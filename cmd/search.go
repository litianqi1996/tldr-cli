package cmd

import (
	"errors"
	"fmt"
)

const TIPS = "This page doesn't exist yet\nSubmit new pages here: https://github.com/tldr-pages/tldr"

func SearchPages(cmd string, repopath string, language string) (string, error) {

	pgs := []string{"pages"}
	if language != "" {
		pgs = append(pgs, "pages."+language)
	}

	platforms := []string{"common", "linux", "osx", "sunos", "windows"}

	for i := len(pgs) - 1; i >= 0; i-- {
		for _, pl := range platforms {
			cmdpath := fmt.Sprintf("%s/%s/%s/%s.md", repopath, pgs[i], pl, cmd)
			ok, err := PathExists(cmdpath)
			if err != nil {
				return "", nil
			}
			if ok {
				return cmdpath, nil
			}

		}
	}

	return "", errors.New(TIPS)
}
