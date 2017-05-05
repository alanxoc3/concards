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

	d, err := deck.Open("sample.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

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

	termgui.Run(sessionDeck)
}
