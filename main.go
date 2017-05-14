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
	if cfg == nil {
		os.Exit(0)
	}

	//cfg.Print()

	decks, err := deck.OpenDecks(cfg.Files)
	do_err(err)

	writeFunc := func() {
		err := decks.Write()
		do_err(err)
	}

	if cfg.UpdateMode {
		writeFunc()

	} else {
		session, err := gen_session_deck(cfg, decks)
		do_err(err)

		if cfg.Usage == termhelp.VIEWMODE {
			err = termboxgui.TermBoxRun(session, cfg, decks)
			writeFunc()
			do_err(err)
		} else if cfg.Usage == termhelp.EDITMODE {
			session.Sort()
			err = deck.EditDeck(cfg.Editor, session)
			writeFunc()
			do_err(err)
		} else if cfg.Usage == termhelp.PRINTMODE {
			session.Sort()
			fmt.Print(session.ToString())
		}
	}
}

func gen_session_deck(cfg *termhelp.Config, decks deck.DeckControls) (session deck.Deck, err error) {
	// Build Deck
	if cfg.Review {
		session = append(session, decks.FilterReview()...)
	}

	if cfg.Memorize {
		session = append(session, decks.FilterMemorize()...)
	}

	if cfg.Done {
		session = append(session, decks.FilterDone()...)
	}

	// Filter Deck
	if cfg.GroupsEnabled {
		session = session.FilterGroupsAdd(cfg.GroupsSlice)
	}

	session.Shuffle()

	if cfg.NumberEnabled {
		session = session.FilterNumber(cfg.Number)
	}

	if len(session) == 0 {
		return nil, fmt.Errorf("Error: There were no cards that met your demands.")
	}

	return
}
