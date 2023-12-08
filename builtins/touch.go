package builtins

import (
	"fmt"
	"os"
)

func Touch(args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("touch: expects an argument")
	}
	if len(args) > 1 {
		return fmt.Errorf("touch: too many arguments")
	}

	filename := args[0]
	
	file, err := os.Create(filename)		
	if err != nil {
		defer file.Close()
		return err
	} else {
		fmt.Println("touch: created", filename)
	}
	defer file.Close()

	return nil
}
