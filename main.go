package main

import (
	"fmt"
	"os"

	"github.com/alanxoc3/concards-go/deck"
	"github.com/alanxoc3/concards-go/termgui"
	"github.com/alanxoc3/concards-go/termhelp"
)

func main() {
	cfg, err := termhelp.ParseConfig(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cfg.MainScreen)

	d, err := deck.Open("sample.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	d.Print()

	termgui.Run()
}
