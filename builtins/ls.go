package builtins

import (
	"fmt"
	"os"
)

func Ls(args ...string) error {
	directory := "."
	if len(args) > 1 {
		return fmt.Errorf("ls: too many arguments")
	} else if len(args) == 1 {
		directory = args[0]
	}

	files, err := os.ReadDir(directory)
	if err != nil {
		return err
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
	return nil
}
