package file

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/alanxoc3/concards/deck"
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
	if cfg == nil {
		panic("Config was nil when passed to edit function.")
	}

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
	if d.IsEmpty() {
		return fmt.Errorf("Error: The deck is empty.")
	}

	// We need to get information for the top card first.
	curHash, curCard, curMeta := d.Top()
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

	// Change the insert index based on if the current card was removed or not.
	cardIndex := 0
	if curHash == d.TopHash() {
		cardIndex = 1
	}

	// Check if the current card was removed or not.
	for i := deckAfter.Len() - 1; i >= 0; i-- {
		newCard := deckAfter.GetCard(i)
		d.InsertCard(newCard, cardIndex)
		d.AddMetaIfNil(newCard.HashStr(), curMeta)
	}

	return nil
}
