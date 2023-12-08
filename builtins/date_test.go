package builtins_test

import (
	"testing"

	"github.com/galexander77/shell-builtins/builtins"
)

func TestDate(t *testing.T) {
    err := builtins.Date()
    if err != nil {
        t.Errorf("%v", err)
    }

    err = builtins.Date("test")
    if err == nil {
        t.Errorf("")
    }
}