package cmd

import (
	"github.com/fatih/color"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func PrintLogo() {

	color.Cyan(`   __     __       __
  / /_   / /  ____/ /  _____
 / __/  / /  / __  /  / ___/
/ /_   / /  / /_/ /  / /
\__/  /_/   \____/  /_/

`)

}
