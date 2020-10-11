package card

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"crypto/sha256"

	"github.com/alanxoc3/concards/internal"
)

// A card is a list of facts. Usually, but not limited to, Q&A format.
type Card struct {
	file  string
	facts [][]string
}

type CardMap map[internal.Hash]*Card

func isSpecialChar(r rune) bool {
	return r == ':' || r == '\\' || r == '|' || r == '>' || r == '<' || r == '{' || r == '}' || unicode.IsSpace(r)
}

func scanCardSides(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !unicode.IsSpace(r) {
			break
		}
	}

	// Check for pipe.
	if len(data) >= start+1 {
		if data[start] == byte('|') || data[start] == byte('{') || data[start] == byte('}') {
			return start + 1, data[start : start+1], nil
		}
	}

	// Check for colons.
	if len(data) >= start+2 {
		if data[start] == byte(':') && data[start+1] == byte(':') {
			return start + 2, data[start : start+2], nil
		}
	}

	// Parse until next token
	isBackslash := false
	isColon := false
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])

		if !isBackslash {
			if unicode.IsSpace(r) {
				return i, data[start:i], nil
			} else if r == '{' || r == '}' || r == '|' || unicode.IsSpace(r) {
				return i, data[start:i], nil
			} else if isColon && r == ':' {
				return i - 1, data[start : i-1], nil
			} else if r == ':' {
				isColon = true
			} else {
				isColon = false
			}
		} else {
			isColon = false
		}

		isBackslash = r == '\\' && !isBackslash
	}

	// Return the non empty word.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}

	// Request more data.
	return start, nil, nil
}

func escapeSymbols(s string) string {
	isBackslash := false
	retStr := ""
	for _, c := range s {
		if isBackslash {
			if isSpecialChar(c) {
				retStr += "\\" + string(c)
			} else {
				retStr += string(c)
			}
		} else if c != '\\' {
			if isSpecialChar(c) {
				retStr += "\\" + string(c)
			} else {
            retStr += string(c)
			}
		}

		isBackslash = c == '\\' && !isBackslash
	}
	return retStr
}

// Returns a list of cards, or an empty list if there is an error.
func NewCards(file string, cardStr string) ([]*Card, error) {
	if file == "" {
		return []*Card{}, fmt.Errorf("File not provided.")
	}

	side := []string{}
	sides := [][]string{}
   revSides := [][][]string{}
	cards := []*Card{}

   // Step 1: Scan through string by card words and special tokens.
	scanner := bufio.NewScanner(strings.NewReader(cardStr + " |"))
	scanner.Split(scanCardSides)
	for scanner.Scan() {
		t := scanner.Text()
      if t == "::" || t == "|" {
         if len(side) > 0 {
            sides = append(sides, side)
            side = []string{}
         }

         if t == "::" && len(sides) > 0 {
            revSides = append(revSides, sides)
            sides = [][]string{}
         }
      } else if len(t) > 0 {
			side = append(side, escapeSymbols(t))
		}
	}

   // Step 2: Put any remaining sides to the reverse card structure.
   if len(sides) > 0 {
      revSides = append(revSides, sides)
   }

   // Step 3: Add all the cards/reverse cards.
   if len(revSides) > 1 {
      for ri, rs := range revSides {
         for _, s := range rs {
            facts := [][]string{}
            facts = append(facts, s)
            for rri, rrs := range revSides {
               if rri != ri {
                  for _, ss := range rrs {
                     facts = append(facts, ss)
                  }
               }
            }

            cards = append(cards, &Card{file, facts})
         }
      }
   } else if len(revSides) == 1 {
      cards = append(cards, &Card{file, revSides[0]})
   }

   return cards, nil
}

func (c *Card) HasAnswer() bool { return len(c.facts) > 1 }
func (c *Card) String() string  { return strings.Join(c.getFactsRaw(), " | ") }
func (c *Card) File() string    { return c.file }
func (c *Card) Len() int        { return len(c.facts) }

func (c *Card) Hash() (dest internal.Hash) {
	hash := sha256.Sum256([]byte(c.String()))
	copy(dest[:], hash[:])
	return dest
}

func (c *Card) GetFactEsc(i int) string {
	factStr := c.getFactRaw(i)

	scanner := bufio.NewScanner(strings.NewReader(factStr))
	scanner.Split(bufio.ScanRunes)
	prev := ""
	str := ""
	for scanner.Scan() {
		t := scanner.Text()

		if prev == "\\" {
			str += t
		} else if t != "\\" {
			str += t
		}
		prev = t
	}
	return str
}

// -------------------- Private stuff below --------------------

func (c *Card) getFactRaw(i int) string {
	return c.getFactHelper(i, func(words []string) []string {
		return words
	})
}

func (c *Card) getFactsRaw() []string {
	return c.getFactsHelper((*Card).getFactRaw)
}

func (c *Card) getFactsHelper(factFunc func(*Card, int) string) []string {
	facts := []string{}
	for i := range c.facts {
		facts = append(facts, factFunc(c, i))
	}
	return facts
}

func (c *Card) getFactHelper(i int, factLogic func([]string) []string) string {
	words := []string{}
	if len(c.facts) > i && 0 <= i {
		words = factLogic(c.facts[i])
	}
	return strings.Join(words, " ")
}

func createReverseCard(file string, facts [][]string) *Card {
	return &Card{file, [][]string{facts[len(facts)-1], facts[0]}}
}
