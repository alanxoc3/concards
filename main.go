package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	// "github.com/alanxoc3/concards/deck"
	// "github.com/alanxoc3/concards/termboxgui"
	"github.com/alanxoc3/concards/termhelp"
	"github.com/alanxoc3/concards/file"
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

	deck, err := file.ReadToDeck(cfg.Files[0])
	do_err(err)

	writeFunc := func() {
		// err := decks.Write()
		// do_err(err)
	}

	if cfg.UpdateMode {
		writeFunc()

	} else {
		if len(deck) > 0 {
			if cfg.Usage == termhelp.VIEWMODE {
				// err := termboxgui.TermBoxRun(session, cfg, decks)
				// writeFunc()
				// do_err(err)
			} else if cfg.Usage == termhelp.EDITMODE {
				deck.Sort()
				// err := deck.EditDeck(cfg.Editor, deck)
				writeFunc()
				do_err(err)
			} else if cfg.Usage == termhelp.PRINTMODE {
				deck.Sort()
				fmt.Print(file.WriteDeckToString(&deck))
			}
		}
	}
}

// func gen_session_deck(cfg *termhelp.Config, decks deck.DeckControls) (session deck.Deck) {
	// Build Deck
	// if cfg.Review {
		// session = append(session, decks.FilterReview()...)
	// }

	// if cfg.Memorize {
		// session = append(session, decks.FilterMemorize()...)
	// }

	// if cfg.Done {
		// session = append(session, decks.FilterDone()...)
	// }

	// Filter Deck
	// if cfg.GroupsEnabled {
		// session = session.FilterGroupsAdd(cfg.GroupsSlice)
	// }

	// session.Shuffle()

	// if cfg.NumberEnabled {
		// session = session.FilterNumber(cfg.Number)
	// }

	// return
// }
