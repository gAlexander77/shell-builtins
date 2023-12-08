package builtins

import (
	"fmt"
	"os"
)

func PrintWorkingDirectory() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(dir)
	return nil
}
