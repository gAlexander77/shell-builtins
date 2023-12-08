package builtins_test

import (
	"testing"

	"github.com/galexander77/shell-builtins/builtins"
)

func TestLs(t *testing.T) {
    err := builtins.Ls()
    if err != nil {
        t.Errorf("%v", err)
    }
    
    err = builtins.Ls("arg1", "arg2")
    if err == nil {
        t.Errorf("")
    }
}