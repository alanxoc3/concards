package deck

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/alanxoc3/concards-go/card"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func EditDeck(editor string, d Deck, message string) {
	env := editor
	d.Sort()

	if env == "" {
		if env = os.Getenv("EDITOR"); env == "" {
			fmt.Println("Error: Your \"EDITOR\" environment variable isn't set.")
			return
		}
	}

	if _, err := exec.LookPath(env); err != nil {
		fmt.Printf("Error: \"%s\" is not installed on this machine.", editor)
		return
	}

	tempFile, err := ioutil.TempFile("", "concards")
	if err != nil {
		fmt.Printf("Error: Couldn't create a temporary file for editing.\n")
		return
	}

	// It doesn't really matter if there is an error removing the temp file.
	defer os.Remove(tempFile.Name())

	err = WriteDeck(&d, tempFile.Name(), message)
	if err != nil {
		fmt.Printf("Error: Couldn't write to a temporary file for editing.\n")
		return
	}

	cmd := exec.Command(env, tempFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error: The editor returned an error code.")
		return
	}

	if dc, err := Open(tempFile.Name()); err != nil {
		fmt.Printf("Error: \"%s\"", err)
		return
	} else {
		copyDeckContents(&d, &dc.Deck)
	}
	// fmt.Printf("%s is the filename.\n", tempFile.Name())
}

func copyDeckContents(dst, src *Deck) {
	for i := 0; i < min(len(*src), len(*dst)); i++ {
		tmpId := (*dst)[i].Id
		*(*dst)[i] = *(*src)[i]
		(*dst)[i].Id = tmpId
	}
}

func EditCard(editor string, c *card.Card, message string) {
	var d Deck
	d = append(d, c)
	EditDeck(editor, d, message)
}
