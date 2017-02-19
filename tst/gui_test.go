package gui

import "fmt"
import "testing"
import "gui"

func TestGui(t *testing.T) {
	cfg := GenConfig()
	// options := cfg.Opts

	fmt.Printf("%s", cfg.Help())
}
