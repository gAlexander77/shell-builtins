package builtins

import (
	"fmt"
	"os"
)

func Rm(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("rm: expects at least one argument")
	} else if len(args) > 1 {
		return fmt.Errorf("rm: too many arguments")
	}

	filename := args[0]
	err := os.Remove(filename)
	if err != nil {
		return fmt.Errorf("rm: unable to remove \"%s\"\nrm: %v", filename, err)
	} else {
		fmt.Println("rm: removed", filename)
	}
	return nil
}
