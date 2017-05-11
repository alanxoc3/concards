package deck

import (
	"io/ioutil"
)

// Assumes a sorted deck by Id.
// Will write in blocks. The first block will be up until a different file is
// found.
func WriteDeckControl(d *DeckControl) error {
	str := []byte(d.ToString())
	err := ioutil.WriteFile(d.Filename, str, 0644)
	if err != nil {
		return err
	}

	return nil
}

func WriteDeck(d *Deck, filename string, message string) error {
	str := []byte(message + "\n\n" + d.ToString())
	err := ioutil.WriteFile(filename, str, 0644)
	if err != nil {
		return err
	}

	return nil
}
