package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"

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
	facts := ""
	state := false
	prev := ""
	var td []*card.Card

	// Scan by words.
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		t := scanner.Text()

		if prev == "\\" {
			prev = "\\" + t
		} else if state {
			if prev == "@" && t == ">" {
				cards, _ := card.NewCards(f, facts)
				td = append(td, cards...)

				facts = ""
			} else if prev == "<" && t == "@" {
				cards, _ := card.NewCards(f, facts)
				td = append(td, cards...)
				cl = append(cl, td...)
				state = false
            prev = ""
			} else {
				facts = facts + prev
            prev = t
			}
		} else if prev == "@" && t == ">" {
			// create td
			td = []*card.Card{}
			state = true
			facts = ""
			prev = ""
		} else {
			prev = t
		}
	}

	return cl
}
