package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/alanxoc3/concards/card"
	"github.com/alanxoc3/concards/deck"
)

// Open opens filename and loads cards into new deck
func ReadCardsToDeck(d *deck.Deck, filename string) error {
	err := filepath.Walk(filename, func(path string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		n := info.Name()
		isHidden := len(n) > 1 && string(n[0]) == "." && n != ".."
		isDir := info.IsDir()

		if isDir && isHidden {
			return filepath.SkipDir
		} else if isHidden || isDir {
			return nil
		}

		absPath, _ := filepath.Abs(path)
		if f, fe := os.Open(absPath); fe != nil {
			return fmt.Errorf("Error: Unable to open file \"%s\"", filename)
		} else {
			defer f.Close()
			ReadCardsToDeckHelper(f, d, absPath)
		}

		return nil
	})

	return err
}

func ReadCardsToDeckHelper(r io.Reader, d *deck.Deck, f string) {
	// Initialization.
	facts := []string{}
	state := false
	var td *deck.Deck

	// Scan by words.
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		t := scanner.Text()

		if state {
			if t == card.CBeg {
				td.AddNewCards(f, strings.Join(facts, " "))

				facts = []string{}
			} else if t == card.CEnd {
				td.AddNewCards(f, strings.Join(facts, " "))

				for i := 0; i < td.Len(); i++ {
					d.AddCard(td.GetCard(i))
				}
				state = false
			} else {
				facts = append(facts, t)
			}
		} else if t == card.CBeg {
			// create td
			td = deck.NewDeck()
			state = true
			facts = []string{}
		}
	}

	return
}
