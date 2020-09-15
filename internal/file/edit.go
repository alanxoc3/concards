package file

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/alanxoc3/concards/internal"
	"github.com/alanxoc3/concards/internal/deck"
)

type DeckFunc func(string, *Config) (*deck.Deck, error)

func ReadCards(filename string, cfg *Config) (*deck.Deck, error) {
	d := deck.NewDeck()
	if err := ReadCardsToDeck(d, filename); err != nil {
		return nil, err
	}
	return d, nil
}

func EditCards(filename string, cfg *Config) (*deck.Deck, error) {
	internal.AssertLogic(cfg != nil, "config was nil when passed to edit function")

	// Load the file with your favorite editor.
	cmd := exec.Command(cfg.Editor, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("Error: The editor returned an error code.")
	}

	return ReadCards(filename, cfg)
}

// Assumes the deck is sorted how you want it to be sorted.
func EditFile(d *deck.Deck, cfg *Config, rf DeckFunc, ef DeckFunc) error {
	if d.ReviewLen() == 0 {
		return fmt.Errorf("Error: The deck is empty.")
	}

	// We need to get information for the top card first.
   curHash := d.TopHash()
   curCard := d.TopCard()
   curMeta := d.TopPredict()
   internal.AssertLogic(curHash != nil && curCard != nil && curMeta != nil, "no top info for non empty deck")

	filename := curCard.File()

	// Deck before editing.
	deckBefore, e := rf(filename, cfg)
	if e != nil {
		return e
	}

	// Deck after editing.
	deckAfter, e := ef(filename, cfg)
	if e != nil {
		return e
	}

	// Take out any card that was removed from the file.
	d.FileIntersection(filename, deckAfter)

	// Get only the cards that were created in the file.
	deckAfter.OuterLeftJoin(deckBefore)
   cl := deckAfter.CardList()

	// Check if the current card was removed or not.
	for i := len(cl) - 1; i >= 0; i-- {
		newCard := cl[i]
      d.AddCards(&newCard)
		d.AddPredicts(curMeta.Clone(newCard.Hash()))
	}

	return nil
}
