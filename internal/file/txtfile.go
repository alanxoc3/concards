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

		cardPath := getPreferredPath(path)
		if f, fe := os.Open(cardPath); fe != nil {
			return fmt.Errorf("Error: Unable to open file \"%s\"", filename)
		} else {
			defer f.Close()
			cl = append(cl, ReadCardsFromReader(f, cardPath)...)
		}

		return nil
	})

	return cl, err
}

func scanCardSections(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if len(data) >= 2 && data[0] == byte(':') && data[1] == byte('#') {
		return 2, data[:2], nil
	}

	// Go to start of first card.
	isBackslash := false
	isHash := false
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])

		if !isBackslash {
			if isHash && r == ':' {
				start += width
				isHash = false
				break
			} else {
				isHash = r == '#'
			}
		} else {
			isHash = false
		}

		isBackslash = r == '\\' && !isBackslash
	}

	isColon := false
	isHash = false
	// Go until start of next card or end of card section.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])

		isBackslash = r == '\\' && !isBackslash
		if isHash && r == ':' {
			return i - 1, data[start : i-1], nil
		} else if isColon && r == '#' {
			return i - 1, data[start : i-1], nil
		} else if !isBackslash && r == ':' {
			isColon = true
			isHash = false
		} else if !isBackslash && r == '#' {
			isHash = true
			isColon = false
		} else {
			isHash = false
			isColon = false
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

		if t == ":#" {
			cl = append(cl, td...)
			td = []*card.Card{}
		} else if len(t) > 0 {
			cards, _ := card.NewCards(f, t)
			td = append(td, cards...)
		}
	}

	return cl
}

// relative path OR absolute path OR default path
func getPreferredPath(defaultPath string) string {
	absPath, err := filepath.Abs(defaultPath)
	if err != nil {
		return defaultPath
	}

	workDir, err := os.Getwd()
	if err != nil {
		return absPath
	}

	prefPath, err := filepath.Rel(workDir, absPath)
	if err != nil {
		return absPath
	}

	return prefPath
}
