package main

import (
	"fmt"
	"os"

	"github.com/alanxoc3/concards-go/deck"
	"github.com/alanxoc3/concards-go/termhelp"
)

func main() {
	_, err := gui.ParseConfig(os.Args)
	if err != nil {
		fmt.Print(err)
	}

	d, err := deck.Open("sample.txt")
	if err != nil {
		fmt.Print(err)
	}

	d.Print()
}
