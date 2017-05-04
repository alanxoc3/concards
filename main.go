package main

import (
	"fmt"
	"os"

	"github.com/alanxoc3/concards-go/deck"
	"github.com/alanxoc3/concards-go/termgui"
	"github.com/alanxoc3/concards-go/termhelp"
)

func main() {
	cfg, err := termhelp.ValidateAndParseConfig(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg.Print()

	if !cfg.Review {
		return
	}

	d, err := deck.Open("sample.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	d.Print()

	fmt.Println("--START--")
	fmt.Print(d.ToStringFromFile("sample.txt"))
	fmt.Println("--END--")

	termgui.Run()
}
