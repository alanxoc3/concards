package main

import (
	"fmt"
	"github.com/alanxoc3/concards-go/gui"
)

func main() {
	cfg := gui.GenConfig()
	// options := cfg.Opts

	fmt.Printf("%s", cfg.Help())
}
