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

// Returns a list of cards, or an empty list if there is an error.
func NewCards(file string, sides string) ([]*Card, error) {
	if file == "" {
		return []*Card{}, fmt.Errorf("File not provided.")
	}

	fact := []string{}
	facts := [][]string{}
	cards := []*Card{}
	waitForReverse := false
	prev := ""
	word := ""

	// Add a separator to the end to not repeat logic.
	sides = sides + " |"
	scanner := bufio.NewScanner(strings.NewReader(sides))
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		t := scanner.Text()

		if prev == "\\" {
			// Only escape special characters.
			if t == ":" || t == "\\" || t == "|" || t == "@" || t == ">" || t == "<" {
				prev = "\\" + t
			} else {
				prev = t
			}
		} else {
			word += prev

			if r, _ := utf8.DecodeRuneInString(t); unicode.IsSpace(r) {
				prev = ""
				if len(word) > 0 {
					fact = append(fact, word)
					word = ""
				}
			} else {
				if t == "|" || t == ":" {
					if len(word) > 0 {
						fact = append(fact, word)
						word = ""
					}

					if len(fact) > 0 {
						facts = append(facts, fact)
						fact = []string{}

						if waitForReverse {
							cards = append(cards, createReverseCard(file, facts))
						}
					}

					waitForReverse = t == ":"
					prev = ""
				} else {
					prev = t
				}
			}
		}
	}

	if len(facts) > 0 {
		return append([]*Card{&Card{file, facts}}, cards...), nil
	} else {
		return []*Card{}, fmt.Errorf("Question not provided.")
	}
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
