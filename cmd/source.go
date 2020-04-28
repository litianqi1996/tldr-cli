package cmd

import (
	"github.com/fatih/color"
	"gopkg.in/src-d/go-git.v4"
	"os"
	"os/signal"
	"syscall"
)

// catch interrupt signal for unhappy you , github is so slow for china sometimes.
func SignalHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		_ = CleanRepo()
		os.Exit(0)
	}()
}

func InitRepo(repoPath string, gitRepo string) error {
	PrintLogo()
	WarningShow("Initialize repository from " + RepoURL)

	SignalHandler()
	_, err := git.PlainClone(repoPath, false, &git.CloneOptions{
		URL:      gitRepo,
		Progress: os.Stdout,
	})

	if err != nil {
		return err
	}
	return err
}

func UpdateRepo() error {
	PrintLogo()
	WarningShow("Update repository from " + RepoURL)

	r, err := git.PlainOpen(RepoPath)
	if err != nil {
		return err
	}
	w, _ := r.Worktree()

	err = w.Pull(&git.PullOptions{
		RemoteName: "origin",
		Progress:   os.Stdout,
		Force:      true,
	})
	if err != nil {
		return err
	}
	err = UpdateTime()
	if err != nil {
		return err
	}
	return err
}

func CleanRepo() error {
	return os.RemoveAll(RepoPath)
}

var WarningShow = color.New(color.FgYellow).PrintlnFunc()
