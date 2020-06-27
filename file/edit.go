package file

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/alanxoc3/concards/core"
)

type DeckFunc func(string, *Config) (*core.Deck, error)

func ReadCards(filename string, cfg *Config) (*core.Deck, error) {
	d := core.NewDeck()
	if err := ReadCardsToDeck(d, filename, cfg.IsSides); err != nil {
		return nil, err
	}
	return d, nil
}

func EditCards(filename string, cfg *Config) (*core.Deck, error) {
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
func EditFile(d *core.Deck, cfg *Config, rf DeckFunc, ef DeckFunc) error {
	if d.IsEmpty() {
		return fmt.Errorf("Error: The deck is empty.")
	}

	// We need to get information for the top card first.
	cur_hash, cur_card, cur_meta := d.Top()
	filename := cur_card.GetFile()

	// Deck before editing.
	deck_before, e := rf(filename, cfg)
	if e != nil {
		return e
	}

	// Deck after editing.
	deck_after, e := ef(filename, cfg)
	if e != nil {
		return e
	}

	// Take out any card that was removed from the file.
	d.FileIntersection(filename, deck_after)

	// Get only the cards that were created in the file.
	deck_after.OuterLeftJoin(deck_before)

	// Change the insert index based on if the current card was removed or not.
	card_index := 0
	if cur_hash == d.TopHash() {
		card_index = 1
	}

	// Check if the current card was removed or not.
	for i := deck_after.Len() - 1; i >= 0; i-- {
		new_card := deck_after.GetCard(i)
		d.InsertCard(new_card, card_index)
		d.AddMetaIfNil(new_card.HashStr(), cur_meta)
	}

	return nil
}
