package deck

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/alanxoc3/concards-go/card"
)

// Deck contains a set of cards
type Deck struct {
	Cards []card.Card
}

type block struct {
	start int
	end   int
	lines []string
}

// Size returns the number of cards in the deck
func (d *Deck) Size() int {
	return len(d.Cards)
}

// Open opens filename and loads cards into new deck
func Open(filename string) (d *Deck, err error) {
	file, err := os.Open(filename)
	if err != nil {
		err = fmt.Errorf("Unable to open deck: %s", err)
		return
	}

	currLine := 0
	scanner := bufio.NewScanner(file)
	for {
		eof, b := nextBlock(scanner, currLine)
		currLine = b.end + 1

		// handle block

		if eof {
			break
		}
	}

	if err = scanner.Err(); err != nil {
		err = fmt.Errorf("Error reading file: %s", err)
	}

	d = &Deck{}
	return
}

func nextBlock(scanner *bufio.Scanner, currLine int) (atEOF bool, b block) {
	currLine++
	b.start = currLine
	b.lines = make([]string, 0)

	scanner.Scan()
	for {
		t := scanner.Text()
		t = strings.TrimRight(t, "\t\n\r ")

		if !scanner.Scan() {
			atEOF = true
		}

		if t == "" {
			break
		} else {
			b.lines = append(b.lines, t)
			currLine++
		}
	}
	b.end = currLine - 1

	return
}
