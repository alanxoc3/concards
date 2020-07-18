package core

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"

	"crypto/sha256"
)

// All the keywords that concards treats special.
const CEsc = "\\"
const CSep = "|"
const CBeg = "@>"
const CEnd = "<@"

var keyWords = map[string]bool {
   CSep: true,
   CBeg: true,
   CEnd: true,
}

// A card is a list of facts. Usually, but not limited to, Q&A format.
type Card struct {
	file  string
	facts [][]string
}

func parseByWords(s string, wordFunc func(string)) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
      wordFunc(scanner.Text())
	}
}

func NewCard(file string, sides string) (*Card, error) {
	fact := []string{}
	facts := [][]string{}

   parseByWords(sides, func(word string) {
      if word == CSep {
         if len(fact) > 0 {
            facts = append(facts, fact)
            fact = []string{}
         }
      } else if len(word) > 0 {
         fact = append(fact, word)
      }
   })

	if len(fact) > 0 {
		facts = append(facts, fact)
	}

	if len(facts) > 0 {
		return &Card{file, facts}, nil
	} else {
		return nil, fmt.Errorf("Question not provided.")
	}
}

func (c *Card) GetSubCards() []*Card {
	subCards := []*Card{}
	question := c.GetFactRaw(0)
	answers := c.GetFactsRaw()[1:]
	for _, answer := range answers {
		if sc, err := NewCard(c.file, fmt.Sprintf("%s %s %s", answer, CSep, question)); err == nil {
			subCards = append(subCards, sc)
		} else {
			panic("Error: Sub card was not created due to bad parent card. This is a logic error and should be fixed.")
		}
	}
	return subCards
}

func (c *Card) HasAnswer() bool {
	return len(c.facts) > 1
}

func (c *Card) String() string {
	return strings.Join(c.GetFactsRaw(), " " + CSep + " ")
}

func (c *Card) Hash() [sha256.Size]byte {
	return sha256.Sum256([]byte(c.String()))
}

func (c *Card) HashStr() string {
	return fmt.Sprintf("%x", c.Hash())[:32]
}

func (c *Card) Len() int {
	return len(c.facts)
}

func (c *Card) getFactHelper(i int, factLogic func([]string)[]string) string {
   words := []string{}
	if len(c.facts) > i && 0 <= i {
      words = factLogic(c.facts[i])
	}
   return strings.Join(words, " ")
}

func (c *Card) GetFactRaw(i int) string {
   return c.getFactHelper(i, func(words []string) []string {
      return words
   })
}

// This banks on the fact that the backslash is an ASCII character at the beginning.
// If the escape character wasn't ASCII, the logic here would have to change.
func (c *Card) GetFactEsc(i int) string {
   return c.getFactHelper(i, func(words []string) []string {
      re := regexp.MustCompile(`^\\+`)
      newWords := []string{}
      for _, word := range words {
         if escStr := re.ReplaceAllString(word, ""); len(escStr) < len(word) && keyWords[escStr] {
            word = word[1:]
         }
         newWords = append(newWords, word)
      }
      return newWords
   })
}

func (c *Card) getFactsHelper(factFunc func(*Card, int)string) []string {
	facts := []string{}
	for i := range c.facts {
		facts = append(facts, factFunc(c,i))
	}
	return facts
}

func (c *Card) GetFactsRaw() []string {
   return c.getFactsHelper((*Card).GetFactRaw)
}

func (c *Card) GetFactsEsc() []string {
   return c.getFactsHelper((*Card).GetFactEsc)
}

func (c *Card) GetFile() string {
	return c.file
}
