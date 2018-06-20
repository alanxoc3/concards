package deck

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/alanxoc3/concards-go/card"
	"github.com/alanxoc3/concards-go/constring"
)

// Used for the file parsing.
type block struct {
	start int
	end   int
	lines []string
}

// Returns true if we a Byte Order Marker at the beginning of the file.
// TODO: Don't make this just a hotfix, get rid of any marker, not just UTF-8.
func isBOM(bom []byte) bool {
	return bom[0] == 0xef && bom[1] == 0xbb && bom[2] == 0xbf
}

// Open opens filename and loads cards into new deck
func Open(filename string) (d *DeckControl, err error) {
	d = &DeckControl{}

	file, err1 := os.Open(filename)
	if err1 != nil {
		err = fmt.Errorf("Error: Unable to open file \"%s\"", filename)
		return
	}

	// Get rid of UTF-8 encoding.
	bom := make([]byte, 3, 3)

	// Returns an error if fewer than 3 bytes were read.
	io.ReadFull(file, bom[:])
	if !isBOM(bom) {
		file.Seek(0, 0)
	}

	// Set up the line stuff
	curLine := 1
	scanner := bufio.NewScanner(file)
	eof := !scanner.Scan()
	var b *block
	loopOnBatch := false // start out reading the preamble (text before cards).

	// Loop through each batch.
	for !eof {
		if loopOnBatch {
			// Read a batch, aka the current group.
			eof, b = readBatch(scanner, curLine)

			// Put the batch into groups and a file.
			od, err2 := batchToDeck(b, &filename)
			if err2 != nil {
				err = err2
				return
			}

			d.AddDeckWithId(od)

		} else {
			// Read file break.
			eof, b = readFileBreak(scanner, curLine)
			fb := ""

			for i, line := range b.lines {
				fb += line
				if !eof || i < len(b.lines)-1 {
					fb += "\n"
				}
			}

			if fb != "" {
				d.AddFileBreak(fb)
			}
		}

		curLine = b.end
		loopOnBatch = !loopOnBatch
	}

	d.Filename = filename

	return
}

func isLineQuestion(line string) bool {
	return !constring.DoesLineBeginWith(line, "\t") &&
		!constring.DoesLineBeginWith(line, card.SPACES) &&
		!constring.DoesLineBeginWith(line, "~~") &&
		constring.Trim(line) != ""
}

// Takes a batch and puts it into a subdeck. Could have parsing errors.
func batchToDeck(batch *block, filename *string) (*DeckControl, error) {
	deck := &DeckControl{}
	i := 0

	// Step 1: The top lines will be groups in a batch, we will go through these.
	for ; i < len(batch.lines); i++ {
		x := batch.lines[i]

		if constring.DoesLineBeginWith(x, "##") {
			x = constring.TrimLineBegin(x, "##")
			deck.AddGroups(constring.StringToList(x))
		} else if x != "" {
			break
		}
	}

	buff := make([]string, 0)
	onQuestion := true

	// Step 2: The rest of the lines will be cards.
	for ; i < len(batch.lines); i++ {
		x := batch.lines[i]

		// This is when you create a card.
		if !onQuestion && isLineQuestion(x) && len(buff) != 0 {
			c, err := card.New(buff)
			if err != nil {
				return nil, fmt.Errorf("Error: \"%s\" near line %d. %s", *filename, batch.start+i, err)
			}
			c.File = *filename
			c.Groups = deck.Groups
			deck.AddCardWithId(c)

			// Remember to empty the buffer for the next card.
			buff = nil
			onQuestion = true
		}

		if !isLineQuestion(x) {
			onQuestion = false
		}

		// Here we add buffer lines to a card.
		if x != "" && len(buff) == 0 || len(buff) > 0 {
			buff = append(buff, x)
		}
	}

	// And there could be one more card.
	if len(buff) != 0 {
		c, err := card.New(buff)
		if err != nil {
			return nil, fmt.Errorf("Error: \"%s\" near line %d. %s", *filename, batch.start+i, err)
		}
		c.File = *filename
		c.Groups = deck.Groups
		deck.AddCardWithId(c)
	}

	// Finished with the batch!!! Woohoo!
	return deck, nil
}

// This needs scanner to have already scanned something before.
func readBatch(scanner *bufio.Scanner, curLine int) (eof bool, b *block) {
	b = &block{}
	b.start = curLine
	b.lines = make([]string, 0)

	onGroup := true

	for !eof {
		// Some preprocessing on the current line
		t := scanner.Text()
		t = strings.TrimRight(t, "\t\n\r ")

		parsedGroup := constring.DoesLineBeginWith(t, "##")

		// Finish the group header section.
		if onGroup && !parsedGroup && t != "" {
			onGroup = false

			// If the last thing is a group, then you are done.
		} else if !onGroup && parsedGroup {
			break
		} else if t == "##" {
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

// This needs scanner to have already scanned something before.
func readFileBreak(scanner *bufio.Scanner, curLine int) (eof bool, b *block) {
	b = &block{}
	b.start = curLine
	b.lines = make([]string, 0)

	// If line is ##, then skip that line. If line is a group, get out of here.
	t := scanner.Text()
	t = strings.TrimRight(t, "\t\n\r ")
	if t == "##" {
		if !scanner.Scan() {
			eof = true
		}

		curLine++
	}

	for !eof {
		// Some preprocessing on the current line
		t := scanner.Text()
		t = strings.TrimRight(t, "\t\n\r ")

		parsedGroup := constring.DoesLineBeginWith(t, "##") && t != "##"

		if parsedGroup {
			break
		} else {
			b.lines = append(b.lines, t)

			curLine++
			if !scanner.Scan() {
				eof = true
			}
		}
	}

	b.end = curLine

	return
}
