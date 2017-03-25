package card

import (
	"errors"
	"fmt"

	"github.com/alanxoc3/concards-go/constring"
)

// Card represents a single flash card. Contains all
// information pertaining to a card.
type Card struct {
	Question string
	Answer   string
	Groups   []string
	Metadata interface{}
}

// Parses and validates block of text as card.
func New(lines []string) (*Card, error) {
	c := &Card{}

	// Assume we are in question first.
	inQuestion := false
	inAnswer := false
	inMeta := false

	// Helper vars with parsing.
	ansDel := false
	grpDel := false
	metDel := false

	for _, line := range lines {
		// Preformatting for multiple spaces.
		if constring.DoesLineBeginWith(line, "    ") {
			line = "\t" + constring.TrimLineBegin(line, "    ")
		}

		ansDel = constring.DoesLineBeginWith(line, "\t")
		grpDel = constring.DoesLineBeginWith(line, "## ")
		metDel = constring.DoesLineBeginWith(line, "~~ ")

		if grpDel {
			// ERROR: Logic, shouldn't get ## in here.
			return nil, errors.New("logic: found group delim in card")
		}

		// Check at the beginning.
		if !(inMeta || inAnswer || inQuestion) {
			if metDel {
				// ERROR: Not only meta
				return nil, errors.New("file: You can't have meta data on its own!")
			} else if ansDel {
				// ERROR: Not only answer
				return nil, errors.New("file: You can't have an answer on its own!")
			} else {
				inQuestion = true
				c.Question = line
			}

			// Check for question.
		} else if inQuestion {
			if metDel {
				inMeta = true
				inQuestion = false
				c.Metadata = constring.TrimLineBegin(line, "~~ ")
			} else if ansDel {
				inAnswer = true
				inQuestion = false
				c.Answer = constring.TrimLineBegin(line, "\t")
			} else {
				c.Question += line
			}

			// Check for answer.
		} else if inAnswer {
			if metDel {
				inMeta = true
				inAnswer = false
				c.Metadata = constring.TrimLineBegin(line, "~~ ")
			} else if ansDel {
				c.Answer += constring.TrimLineBegin(line, "\t")
			} else {
				// ERROR: Not only meta
				return nil, errors.New("file: Can't have a question in the answer.")
			}

			// This means there was stuff after the meta data.
		} else {
			// assert(inMeta && !inQuestion && !inAnswer)
			return nil, errors.New("file: Found extra lines after the metadata.")
		}
	}

	return c, nil
}

// Prints out the card, for debugging purposes.
func (c *Card) Print() {
	fmt.Println("G: ", c.Groups)
	fmt.Println("Q: ", c.Question)
	fmt.Println("A: ", c.Answer)
	fmt.Println("M: ", c.Metadata)

	fmt.Println()
}
