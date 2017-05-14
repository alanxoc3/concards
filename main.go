package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/alanxoc3/concards-go/deck"
	"github.com/alanxoc3/concards-go/termboxgui"
	"github.com/alanxoc3/concards-go/termhelp"
)

func main() {
	do_err := func(err error) {
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	}

	rand.Seed(time.Now().UTC().UnixNano())

	cfg, err := termhelp.ValidateAndParseConfig(os.Args)
	do_err(err)

	decks, err := deck.OpenDecks(cfg.Files)
	do_err(err)

	var sessionDeck deck.Deck

	if cfg.Review {
		sessionDeck = append(sessionDeck, decks.FilterReview()...)
	}

	if cfg.Memorize {
		sessionDeck = append(sessionDeck, decks.FilterMemorize()...)
	}

	if cfg.Done {
		sessionDeck = append(sessionDeck, decks.FilterDone()...)
	}

	sessionDeck.Shuffle()

	err = termboxgui.TermBoxRun(sessionDeck, cfg, decks)
	do_err(err)

	// For writing the deck
	err = deck.WriteDeckControls(decks)
	do_err(err)
}
