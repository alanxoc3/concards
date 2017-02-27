package card

import (
	"errors"
	"fmt"
)

// Card represents a single flash card. Contains all
// information pertaining to a card.
type Card struct {
	Question string
	Answer   string
	Groups   []string
	Metadata interface{}
}

// New parses and validates block of text as card
func New(lines []string) (c *Card, err error) {
	c = &Card{}
	inAnswer := false
	inQuestion := false
	cardDone := false

	for _, line := range lines {
		if cardDone {
			err = errors.New("Found extra lines after card was finished")
		}

		lineRunes := []rune(line)
		if !inQuestion && !inAnswer && lineRunes[0] != '\t' {
			inQuestion = true
			c.Question = line
		} else if inQuestion && lineRunes[0] != '\t' {
			c.Question = fmt.Sprintf("%s\n%s", c.Question, line)
		} else if inQuestion && lineRunes[0] == '\t' {
			inQuestion = false
			inAnswer = true
			c.Answer = line
		} else if inAnswer && lineRunes[0] == '\t' {
			c.Answer = fmt.Sprintf("%s\n%s", c.Answer, line)
		} else if inAnswer && lineRunes[0] != '\t' && lineRunes[0] != '~' {
			err = errors.New("Cannot have two questions for one card.\nMissing a newline?")
			return
		} else if inAnswer && lineRunes[0] == '~' && lineRunes[1] == '~' {
			c.Metadata = line
			cardDone = true
		}
	}

	return
}
