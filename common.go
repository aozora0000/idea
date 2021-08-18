package idea

import (
	"fmt"
	"os"
	"path/filepath"
)

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func MakeShelf() error {
	if Exists(getShelfDir()) {
		return nil
	}
	return os.Mkdir(getShelfDir(), os.ModePerm)
}

func getShelfDir() string {
	conf, err := os.UserConfigDir()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return filepath.Join(conf, "idea")
}
