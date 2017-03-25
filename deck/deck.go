package deck

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/alanxoc3/concards-go/card"
)

// Deck contains a set of cards
type Deck struct {
	Cards []*card.Card
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
	d = &Deck{}

	file, err := os.Open(filename)
	if err != nil {
		err = fmt.Errorf("Unable to open deck: %s", err)
		return
	}

	currLine := 0
	currGroup := ""
	scanner := bufio.NewScanner(file)
	scanner.Scan() // required for nextBlock() to work cleanly
	for {
		eof, b := nextBlock(scanner, currLine)
		currLine = b.end + 1

		if len(b.lines) > 0 {
			if group := pullGroup(b); group != "" {
				currGroup = group
			} else {
				var c *card.Card
				c, err = card.New(b.lines)
				if err != nil {
					err = fmt.Errorf("Unable to load card starting at %s:%d: %s",
						filename, b.start, err)
					return
				}

				if currGroup != "" {
					c.Groups = append(c.Groups, currGroup)
				}

				d.Cards = append(d.Cards, c)
			}
		}

		if eof {
			break
		}
	}

	if err = scanner.Err(); err != nil {
		err = fmt.Errorf("Error reading file: %s", err)
	}

	return
}

var groupRegex = regexp.MustCompile(`##\s+(\w[\s\w]*)`)

func pullGroup(b block) (grp string) {
	for _, line := range b.lines {
		if groupRegex.MatchString(line) {
			grp = groupRegex.FindStringSubmatch(line)[1]
		}
	}

	return
}

func nextBlock(scanner *bufio.Scanner, currLine int) (atEOF bool, b block) {
	currLine++
	b.start = currLine
	b.lines = make([]string, 0)

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

// Prints out the cards in the deck, for debugging purposes.
func (d *Deck) Print() {
	for i, c := range d.Cards {
		fmt.Printf("Card %d\n", i)
		c.Print()
	}
}
