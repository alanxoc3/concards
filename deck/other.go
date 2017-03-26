package deck

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/alanxoc3/concards-go/card"
	"github.com/alanxoc3/concards-go/constring"
)

// Contains the group divisions, as parsed in the file.
type SubDeck struct {
	Cards  []*card.Card
	groups []*string
}

// Deck contains a set of groupDecks
type Deck struct {
	SubDecks []*SubDeck
}

// Used for the file parsing.
type block struct {
	start int
	end   int
	lines []string
}

// Size for a subdeck
func (d *SubDeck) Size() int {
	return len(d.Cards)
}

// Size returns the number of cards in the deck
func (d *Deck) Size() (total int) {
	for _, x := range d.SubDecks {
		total += x.Size()
	}
	return
}

// Open opens filename and loads cards into new deck
func Open(filename string) (d *Deck, err error) {
	d = &Deck{}

	file, err := os.Open(filename)
	if err != nil {
		err = fmt.Errorf("Unable to open deck: %s", err)
		return
	}

	// Set up the line stuff
	curLine := 1
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	eof := false
	var b *block

	// Loop through each batch.
	for !eof {
		// Step 1: Read a batch, aka the current group.
		eof, b = readBatch(scanner, curLine)
		curLine = b.end

		// Step 2: Debug the batch.
		fmt.Println("CurLine: ", b.start)
		for _, x := range b.lines {
			fmt.Println(x)
		}
	}
	return
}

// This needs scanner to have already scanned something before.
func readBatch(scanner *bufio.Scanner, curLine int) (eof bool, b *block) {
	b = &block{}
	b.start = curLine
	b.lines = make([]string, 0)

	onGroup := true

	for !eof {
		t := scanner.Text()
		t = strings.TrimRight(t, "\t\n\r ")

		parsedGroup := constring.DoesLineBeginWith(t, "## ")

		// Finish the group header section.
		if onGroup && !parsedGroup {
			onGroup = false

			// If the last thing is a group, then you are done.
		} else if !onGroup && parsedGroup {
			break
		}

		// Will add a line if we are still looking for lines.
		b.lines = append(b.lines, t)

		if !scanner.Scan() {
			eof = true
		}

		curLine++
	}

	b.end = curLine

	return
}

// Prints out the cards in the deck, for debugging purposes.
func (d *Deck) Print() {
	count := 0
	for _, sd := range d.SubDecks {
		for _, c := range sd.Cards {
			count += 1
			fmt.Printf("Card %d\n", count)
			c.Print()
		}
	}
}
