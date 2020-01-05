package deck

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/alanxoc3/concards/card"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// Assumes the deck is sorted how you want it to be sorted.
func EditDeck(editor string, d Deck) error {
	message := "You may ONLY EDIT the cards here.\nREARRANGING, DELETING, or ADDING cards WILL CORRUPT your files."

	env := editor

	// TODO: If I'm assuming a sort, why am I calling sort here?
	d.Sort()

	if env == "" {
		if env = os.Getenv("EDITOR"); env == "" {
			return fmt.Errorf("Error: Your \"EDITOR\" environment variable isn't set.")
		}
	}

	if _, err := exec.LookPath(env); err != nil {
		return fmt.Errorf("Error: \"%s\" is not installed on this machine.", editor)
	}

	tempFile, err := ioutil.TempFile("", "concards")
	if err != nil {
		return fmt.Errorf("Error: Couldn't create a temporary file for editing.\n")
	}

	// It doesn't really matter if there is an error removing the temp file.
	defer os.Remove(tempFile.Name())

	err = WriteDeck(&d, tempFile.Name(), message)
	if err != nil {
		return fmt.Errorf("Error: Couldn't write to a temporary file for editing.\n")
	}

	cmd := exec.Command(env, tempFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Error: The editor returned an error code.")
	}

	if dc, err := OpenNewFormat(tempFile.Name()); err != nil {
		return err
	} else {
		copyDeckContents(&d, &dc.Deck)
	}

	return nil
}

func copyDeckContents(dst, src *Deck) {
	for i := 0; i < min(len(*src), len(*dst)); i++ {
		tmpId := (*dst)[i].Id
		*(*dst)[i] = *(*src)[i]
		(*dst)[i].Id = tmpId
	}
}

func EditCard(editor string, c *card.Card) error {
	var d Deck
	d = append(d, c)
	return EditDeck(editor, d)
}
