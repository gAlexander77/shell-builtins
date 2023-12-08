package builtins

import (
	"fmt"
	"time"
)

func Date(args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("date: too many arguments")
	}

	currentTime := time.Now()
	fmt.Println(currentTime.Format(time.UnixDate))
	return nil
}
