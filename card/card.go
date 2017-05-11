package card

import (
	"errors"
	"fmt"

	"github.com/alanxoc3/concards-go/algs"
	"github.com/alanxoc3/concards-go/constring"
)

// Card represents a single flash card. Contains all
// information pertaining to a card.
type Card struct {
	Question string
	Answer   string
	Groups   []string
	Metadata algs.SpaceAlg
	Id       int
	File     string
}

// Parses and validates block of text as card.
func New(lines []string) (*Card, error) {
	c := &Card{}

	alg, _ := algs.New("")
	c.Metadata = *alg

	// Assume we are in question first.
	inQuestion := false
	inAnswer := false
	inMeta := false
	inEmpty := false

	// Helper vars with parsing.
	ansDel := false
	grpDel := false
	metDel := false
	empDel := false

	for _, line := range lines {
		// Preformatting for multiple spaces.
		if constring.DoesLineBeginWith(line, "    ") {
			line = "\t" + constring.TrimLineBegin(line, "    ")
		}

		ansDel = constring.DoesLineBeginWith(line, "\t")
		grpDel = constring.DoesLineBeginWith(line, "## ")
		metDel = constring.DoesLineBeginWith(line, "~~ ")
		empDel = (constring.Trim(line) == "")

		if grpDel {
			// ERROR: Logic, shouldn't get ## in here.
			return nil, errors.New("Logic Error: Found group delim in card")
		}

		// Check at the beginning.
		if !(inMeta || inAnswer || inQuestion) {
			if metDel {
				// ERROR: Not only meta
				return nil, errors.New("You can't have meta data on its own!")
			} else if ansDel {
				// ERROR: Not only answer
				return nil, errors.New("You can't have an answer on its own!")
			} else if !empDel {
				inQuestion = true
				c.Question = line
			}

			// Check for question.
		} else if inQuestion {
			if metDel {
				inMeta = true
				inQuestion = false

				alg, err := algs.New(constring.TrimLineBegin(line, "~~ "))
				if err != nil {
					return nil, err
				}
				c.Metadata = *alg
			} else if ansDel {
				inAnswer = true
				inQuestion = false
				c.Answer = constring.TrimLineBegin(line, "\t")
			} else if !empDel {
				c.Question += "\n" + line
			}

			// Check for answer.
		} else if inAnswer {
			if metDel {
				inMeta = true
				inAnswer = false

				alg, err := algs.New(constring.TrimLineBegin(line, "~~ "))
				if err != nil {
					return nil, err
				}
				c.Metadata = *alg
			} else if ansDel {
				// Add a single new line for all the empty lines in the answer.
				if inEmpty {
					c.Answer += "\n\t"
					inEmpty = false
				}

				c.Answer += "\n" + constring.TrimLineBegin(line, "\t")
			} else if empDel {
				inEmpty = true
			} else {
				// ERROR: Not only meta
				return nil, errors.New("Can't have a question in the answer.")
			}

			// This means there was stuff after the meta data.
		} else if !empDel {
			// assert(inMeta && !inQuestion && !inAnswer)
			return nil, errors.New("Found extra lines after the metadata.")
		}
	}

	return c, nil
}

// Prints out the card, for debugging purposes.
func (c *Card) Print() {
	fmt.Print("G: ")
	for _, x := range c.Groups {
		fmt.Printf("%s, ", x)
	}
	fmt.Println()

	fmt.Println("#: ", c.Id)
	fmt.Println("F: ", c.File)
	fmt.Println("Q: ", c.Question)
	fmt.Println("A: ", c.Answer)
	fmt.Println("M: ", c.Metadata)

	fmt.Println()
}

// Prints out the card according to the format needed for the file.
func (c *Card) FormatQuestion() string {
	str := c.Question
	return str
}

// Prints out the card according to the format needed for the file.
func (c *Card) FormatAnswer() string {
	str := "\t" + constring.TabsToNewlines(c.Answer)
	return str
}

// Prints out the card according to the format needed for the file.
func (c *Card) FormatGroups() string {
	str := constring.GroupListToString(c.Groups)
	return str
}

// Prints out the card according to the format needed for the file.
func (c *Card) FormatFile() string {
	str := c.Question

	if c.Answer != "" {
		str += "\n\t" + constring.TabsToNewlines(c.Answer)
	}

	str += "\n~~ " + c.Metadata.ToString()

	return str
}
