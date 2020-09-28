package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/alanxoc3/concards/internal"
	"github.com/alanxoc3/concards/internal/card"
)

// Open opens filename and loads cards into new deck
func ReadCardsFromFile(filename string) ([]*card.Card, error) {
	cl := []*card.Card{}
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
			cl = append(cl, ReadCardsFromReader(f, absPath)...)
		}

		return nil
	})

	return cl, err
}

func ReadCardsFromReader(r io.Reader, f string) []*card.Card {
	// Initialization.
	cl := []*card.Card{}
	facts := []string{}
	state := false
	var td []*card.Card

	// Scan by words.
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		t := scanner.Text()

		if state {
			if t == internal.CBeg {
				cards, _ := card.NewCards(f, strings.Join(facts, " "))
				td = append(td, cards...)

				facts = []string{}
			} else if t == internal.CEnd {
				cards, _ := card.NewCards(f, strings.Join(facts, " "))
				td = append(td, cards...)
				cl = append(cl, td...)
				state = false
			} else {
				facts = append(facts, t)
			}
		} else if t == internal.CBeg {
			// create td
			td = []*card.Card{}
			state = true
			facts = []string{}
		}
	}

	return cl
}
