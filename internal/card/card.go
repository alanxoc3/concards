package card

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"

	"crypto/sha256"

	"github.com/alanxoc3/concards/internal"
)

// A card is a list of facts. Usually, but not limited to, Q&A format.
type Card struct {
	file  string
	facts [][]string
}

// Returns a list of cards, or an empty list if there is an error.
func NewCards(file string, sides string) ([]*Card, error) {
	if file == "" {
		return []*Card{}, fmt.Errorf("File not provided.")
	}

	fact := []string{}
	facts := [][]string{}
	cards := []*Card{}
	waitForReverse := false

	// Add a separator to the end to not repeat logic.
	sides = sides + " " + internal.CSep

	parseByWords(sides, func(word string) {
		if word == internal.CSep || word == internal.CRev {
			if len(fact) > 0 {
				facts = append(facts, fact)
				fact = []string{}

				if waitForReverse {
					cards = append(cards, createReverseCard(file, facts))
				}
			}

			waitForReverse = word == internal.CRev
		} else if len(word) > 0 {
			fact = append(fact, word)
		}
	})

	if len(facts) > 0 {
		return append([]*Card{&Card{file, facts}}, cards...), nil
	} else {
		return []*Card{}, fmt.Errorf("Question not provided.")
	}
}

func (c *Card) HasAnswer() bool { return len(c.facts) > 1 }
func (c *Card) String() string  { return strings.Join(c.getFactsRaw(), " "+internal.CSep+" ") }
func (c *Card) File() string    { return c.file }
func (c *Card) Len() int        { return len(c.facts) }

func (c *Card) Hash() (dest internal.Hash) {
	hash := sha256.Sum256([]byte(c.String()))
	copy(dest[:], hash[:])
	return dest
}

func (c *Card) GetFactEsc(i int) string {
	// This banks on the fact that the backslash is an ASCII character at the beginning.
	// If the escape character wasn't ASCII, the logic here would have to change.
	return c.getFactHelper(i, func(words []string) []string {
		re := regexp.MustCompile(`^\\+`)
		newWords := []string{}
		for _, word := range words {
			if escStr := re.ReplaceAllString(word, ""); len(escStr) < len(word) && internal.KeyWords[escStr] {
				word = word[1:]
			}
			newWords = append(newWords, word)
		}
		return newWords
	})
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

func parseByWords(s string, wordFunc func(string)) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wordFunc(scanner.Text())
	}
}

func createReverseCard(file string, facts [][]string) *Card {
	return &Card{file, [][]string{facts[len(facts)-1], facts[0]}}
}
