package core

import (
	"strings"
	"fmt"

   "crypto/sha256"
)

// A card is a list of facts. Usually, but not limited to, Q&A format.
type Card struct {
   file  string
   facts []string
}

func NewCard(facts [][]string) (*Card, error) {
   c := Card{}
	for _, x := range facts {
      if len(x) > 0 {
         c.facts = append(c.facts, strings.Join(x, " "))
      }
	}

   if len(c.facts) > 0 {
      return &c, nil
   } else {
      return nil, fmt.Errorf("Question not provided.")
   }
}

func (c *Card) HasAnswer() bool {
   return len(c.facts) > 1
}

func (c *Card) GetQuestion() string {
   if len(c.facts) > 0 {
      return c.facts[0]
   } else {
      return ""
   }
}

func (c *Card) String() string {
   return strings.Join(c.facts, " @ ")
}

func (c *Card) Hash() [sha256.Size]byte {
   return sha256.Sum256([]byte(c.String()))
}

func (c *Card) HashStr() string {
   return fmt.Sprintf("%x", c.Hash())[:32]
}
