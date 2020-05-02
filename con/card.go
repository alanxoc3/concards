package con

import (
	"strings"
	"fmt"

   "crypto/sha256"
)

// A card is a list of facts. Usually, but not limited to, Q&A format.
type Card []string

func NewCard(facts [][]string) (Card, error) {
   c := Card{}
	for _, x := range facts {
      if len(x) > 0 {
         c = append(c, strings.Join(x, " "))
      }
	}

   if len(c) > 0 {
      return c, nil
   } else {
      return nil, fmt.Errorf("Question not provided.")
   }
}

func (c Card) HasAnswer() bool {
   return len(c) > 1
}

func (c Card) GetQuestion() string {
   if len(c) > 0 {
      return c[0]
   } else {
      return ""
   }
}

func (c Card) KeyText() string {
   return strings.Join(c, " @ ")
}

func (c Card) Hash() [sha256.Size]byte {
   return sha256.Sum256([]byte(c.KeyText()))
}

func (c Card) HashStr() string {
   return fmt.Sprintf("%x", c.Hash())[:32]
}
