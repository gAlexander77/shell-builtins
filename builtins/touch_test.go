package builtins_test

import (
	"testing"

	"github.com/galexander77/shell-builtins/builtins"
)

func TestTouch(t *testing.T) {
	err := builtins.Touch("test.txt", "test2.txt",)
	if err == nil {
		t.Fatal("")
	}
	
	err = builtins.Touch()
	if err == nil {
		t.Errorf("%v", err)
	}
}