package cmd

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

const (
	DefaultRepo     = "https://github.com/tldr-pages/tldr.git"
	DefaultLanguage = ""
	ExpireTime      = 60 * 60 * 24 * 30
)

var (
	TldrPath string = HomePath() + "/.tldrtmp"
	ConfPath string = TldrPath + "/tldr.yaml"
	RepoPath string
	RepoURL  string
	Language string
)

type Config struct {
	GitRepo    string `yaml:"gitrepo"`
	Language   string `yaml:"language"`
	UpdateTime int64  `yaml:"updatetime"`
}

func init() {
	ok, err := PathExists(TldrPath)
	if err != nil {
		log.Println(err)
	}
	if !ok {
		err = os.MkdirAll(TldrPath, 0755)
		if err != nil {
			log.Println(err)
		}
	}
	ok, err = PathExists(ConfPath)
	if err != nil {
		log.Println(err)
	}
	if !ok {
		err = initConf(ConfPath)
		if err != nil {
			log.Println(err)
		}
	}
	return
}

func initConf(path string) error {
	cf := &Config{
		GitRepo:    DefaultRepo,
		Language:   DefaultLanguage,
		UpdateTime: time.Now().Unix(),
	}

	cj, err := yaml.Marshal(cf)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(ConfPath, cj, 0644)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTime() error {
	var cf Config
	rc, err := ioutil.ReadFile(ConfPath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(rc, &cf)
	if err != nil {
		return errors.New("config yaml file invalid")
	}

	cf.UpdateTime = time.Now().Unix()

	cj, err := yaml.Marshal(cf)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(ConfPath, cj, 0644) // work with os.O_TRUNC
	if err != nil {
		return err
	}

	return err
}

func StartUp() error {

	var cf Config

	rc, err := ioutil.ReadFile(ConfPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(rc, &cf)
	if err != nil {
		return errors.New("config yaml file invalid")
	}

	RepoPath = filepath.Join(TldrPath, strings.Split(cf.GitRepo, "/")[3])
	RepoURL = cf.GitRepo
	if CheckLanguage(cf.Language) {
		Language = cf.Language
	} else {
		return errors.New("language config invalid ")
	}
	CheckExpire(cf.UpdateTime)

	ok, err := PathExists(RepoPath)
	if err != nil {
		fmt.Println(err)
	}

	if !ok {
		err = InitRepo(RepoPath, cf.GitRepo)
		if err != nil {
			fmt.Println(err)
		}
	}

	return err
}

func CheckExpire(ut int64) {
	if time.Now().Unix()-ut > ExpireTime {
		color.Yellow("Local data is older than one month, use -u to update it")
	}
}

func CheckLanguage(lg string) bool {
	lgs := []string{"", "de", "es", "fr", "hbs", "it", "ja", "ko", "pt_BR", "pt_PT", "ta", "zh"}
	for _, j := range lgs {
		if j == lg {
			return true
		}
	}
	return false
}

func HomePath() string {
	currentUser, _ := user.Current()
	return currentUser.HomeDir
}
