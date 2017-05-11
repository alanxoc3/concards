package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/alanxoc3/concards-go/deck"
	"github.com/alanxoc3/concards-go/termgui"
	"github.com/alanxoc3/concards-go/termhelp"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	cfg, err := termhelp.ValidateAndParseConfig(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg.Print()

	d, err := deck.Open(cfg.Files[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println("----------- file begin  -----------")
	//fmt.Print(d.ToString())
	//fmt.Println("----------- file end    -----------")

	//fmt.Println("Printing File Breaks...")
	//for _, x := range d.FileBreaks {
	//x.Print()
	//}

	//d.Deck.Print()

	// fmt.Println("--START--")
	// fmt.Print(d.ToStringFromFile("sample.txt"))
	// fmt.Println("--END--")

	d.Deck.Shuffle()

	var sessionDeck deck.Deck

	if cfg.Review {
		sessionDeck = append(sessionDeck, d.Deck.FilterReview()...)
	}

	if cfg.Memorize {
		sessionDeck = append(sessionDeck, d.Deck.FilterMemorize()...)
	}

	if cfg.Done {
		sessionDeck = append(sessionDeck, d.Deck.FilterDone()...)
	}

	termgui.Run(sessionDeck, cfg)

	// For writing the deck
	if err := deck.WriteDeckControl(d); err != nil {
		fmt.Println(err)
		return
	}

}
