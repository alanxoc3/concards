package main

import (
	"fmt"
	"./gui"
)

func main() {
	cfg := gui.GenConfig()
	// options := cfg.Opts

	fmt.Printf("%s", cfg.Help())
}
