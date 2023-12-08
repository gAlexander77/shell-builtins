package builtins_test

import (
	"testing"

	"github.com/galexander77/shell-builtins/builtins"
)

func TestPrintWorkingDirectory(t *testing.T) {
    err := builtins.PrintWorkingDirectory()
    if err != nil {
        t.Errorf("%v", err)
    }
}
