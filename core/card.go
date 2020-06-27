package core

import (
	"bufio"
	"fmt"
	"strings"

	"crypto/sha256"
)

// A card is a list of facts. Usually, but not limited to, Q&A format.
type Card struct {
	file  string
	facts [][]string
}

func NewCard(file string, sides string) (*Card, error) {
	fact := []string{}
	facts := [][]string{}

	scanner := bufio.NewScanner(strings.NewReader(sides))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		t := scanner.Text()
		if t == "@" {
			if len(fact) > 0 {
				facts = append(facts, fact)
				fact = []string{}
			}
		} else if len(t) > 0 {
			fact = append(fact, t)
		}
	}

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
	question := c.GetQuestion()
	answers := c.GetFacts()[1:]
	for _, answer := range answers {
		if sc, err := NewCard(c.file, answer+" @ "+question); err == nil {
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
	return strings.Join(c.GetFacts(), " @ ")
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

func (c *Card) GetFact(i int) string {
	if len(c.facts) > i {
		return strings.Join(c.facts[i], " ")
	} else {
		return ""
	}
}

func (c *Card) GetQuestion() string {
	return c.GetFact(0)
}

func (c *Card) GetFacts() []string {
	facts := []string{}
	for i := range c.facts {
		facts = append(facts, c.GetFact(i))
	}
	return facts
}

func (c *Card) GetFile() string {
	return c.file
}
