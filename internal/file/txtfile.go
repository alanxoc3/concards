package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"unicode/utf8"

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

func scanCardSections(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if len(data) >= 2 && data[0] == byte('<') && data[1] == byte('@') {
		return 2, data[:2], nil
	}

	// Go to start of first card.
	isBackslash := false
	isAt := false
	isLt := false
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])

		if !isBackslash {
			if isAt && r == '>' {
				start += width
				isAt = false
				break
			} else {
				isAt = r == '@'
			}
		} else {
			isAt = false
		}

		isBackslash = r == '\\' && !isBackslash
	}

	// Go until start of next card or end of card section.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])

		isBackslash = r == '\\' && !isBackslash
		if !isBackslash && r == '<' {
			isLt = true
			isAt = false
		} else if isLt && r == '@' {
			return i - 1, data[start : i-1], nil
		} else if !isBackslash && r == '@' {
			isAt = true
		} else if isAt && r == '>' {
			return i - 1, data[start : i-1], nil
		} else {
			isAt = false
			isLt = false
		}
	}

	// If we are at the EOF, there was no end delimiter, so return nothing.
	if atEOF {
		return len(data), nil, nil
	}

	// Request more data.
	return start, nil, nil
}

func ReadCardsFromReader(r io.Reader, f string) []*card.Card {
	// Initialization.
	cl := []*card.Card{}
	var td []*card.Card

	// Scan by card sections.
	scanner := bufio.NewScanner(r)
	scanner.Split(scanCardSections)
	for scanner.Scan() {
		t := scanner.Text()

		if t == "<@" {
			cl = append(cl, td...)
			td = []*card.Card{}
		} else if len(t) > 0 {
			cards, _ := card.NewCards(f, t)
			td = append(td, cards...)
		}
	}

	return cl
}
