package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/alanxoc3/concards/core"
)

// Open opens filename and loads cards into new deck
func ReadCardsToDeck(d *core.Deck, filename string, includeSides bool) error {
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
			ReadCardsToDeckHelper(f, d, absPath, includeSides)
		}

		return nil
	})

	return err
}

func ReadCardsToDeckHelper(r io.Reader, d *core.Deck, f string, includeSides bool) {
	// Initialization.
	facts := []string{}
	state := false
	var td *core.Deck

	// Scan by words.
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		t := scanner.Text()

		if state {
			if t == core.CBeg {
				td.AddCardFromSides(f, strings.Join(facts, " "), includeSides)

				facts = []string{}
			} else if t == core.CEnd {
				td.AddCardFromSides(f, strings.Join(facts, " "), includeSides)

				for i := 0; i < td.Len(); i++ {
					d.AddCard(td.GetCard(i))
				}
				state = false
			} else {
				facts = append(facts, t)
			}
		} else if t == core.CBeg {
			// create td
			td = core.NewDeck()
			state = true
			facts = []string{}
		}
	}

	return
}
