package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/alanxoc3/concards/termhelp"
	"github.com/alanxoc3/concards/termboxgui"
	"github.com/alanxoc3/concards/file"
	"github.com/alanxoc3/concards/deck"
	"github.com/alanxoc3/concards/deckdb"

    "crypto/sha256"
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

   println(cfg.ConfigFolder)
   println(cfg.DatabasePath)
   println(cfg.ConfigFile)

   sum := sha256.Sum256([]byte("THIS_IS_A_QUESTION"))
	fmt.Printf("%x\n", sum)

   deckdb.OpenDb(cfg.DatabasePath)

   cards := deck.Deck{}
   for _, f := range cfg.Files {
      d, err := file.ReadToDeck(f)
      do_err(err)
      cards = append(cards, d...)
   }

	writeFunc := func() {
		// err := decks.Write()
		// do_err(err)
	}

	if cfg.UpdateMode {
		writeFunc()

	} else {
		if len(cards) > 0 {
			if cfg.Usage == termhelp.VIEWMODE {
				err := termboxgui.TermBoxRun(cards, cfg)
				do_err(err)
				// writeFunc()
			} else if cfg.Usage == termhelp.EDITMODE {
				cards.Sort()
				err := file.EditDeck(cfg.Editor, cards)
				do_err(err)
				// writeFunc()
			} else if cfg.Usage == termhelp.PRINTMODE {
				cards.Sort()
				fmt.Print(file.WriteDeckToString(&cards))
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
