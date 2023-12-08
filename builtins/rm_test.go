package builtins_test

import (
	"testing"

	"github.com/galexander77/shell-builtins/builtins"
)

func TestRm(t *testing.T) {
    err := builtins.Rm()
    if err == nil {
        t.Errorf("")
    }

    err = builtins.Rm("file1", "file2")
    if err == nil {
        t.Errorf("")
    }
    
    err = builtins.Rm("file.txt")
    if err == nil {
        t.Errorf("")
    }
}
