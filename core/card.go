package core

import (
	"strings"
	"fmt"

   "crypto/sha256"
)

// A card is a list of facts. Usually, but not limited to, Q&A format.
type Card struct {
   File  string
   Facts []string
}

// Assumes a "cleaned" file string.
func NewCard(facts [][]string, file string) (*Card, error) {
   c := Card{}
   c.File = file
	for _, x := range facts {
      if len(x) > 0 {
         c.Facts = append(c.Facts, strings.Join(x, " "))
      }
	}

   if len(c.Facts) > 0 {
      return &c, nil
   } else {
      return nil, fmt.Errorf("Question not provided.")
   }
}

func (c *Card) HasAnswer() bool {
   return len(c.Facts) > 1
}

func (c *Card) GetQuestion() string {
   if len(c.Facts) > 0 {
      return c.Facts[0]
   } else {
      return ""
   }
}

func (c *Card) String() string {
   return strings.Join(c.Facts, " @ ")
}

func (c *Card) Hash() [sha256.Size]byte {
   return sha256.Sum256([]byte(c.String()))
}

func (c *Card) HashStr() string {
   return fmt.Sprintf("%x", c.Hash())[:32]
}
