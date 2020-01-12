package card

import (
	"strings"
	"sort"

	"github.com/alanxoc3/concards/algs"
)

type Groups map[string]bool

func (g Groups) ToArray() (keys []string) {
   keys = make([]string, 0, len(g))
   for key := range g {
      keys = append(keys, key)
   }
   sort.Strings(keys)
   return
}

const SPACES string = "   "

// Card represents a single flash card. Contains all
// information pertaining to a card.
type Card struct {
	Groups   Groups
	Question string
	Answers  []string
	Notes    []string
	Metadata algs.SpaceAlg
	Deleted  bool
}

func New(
   groups map[string]bool,
   question []string,
   answers [][]string,
   notes [][]string,
   meta []string) (c *Card) {
	c = &Card{}
   c.Question = strings.Join(question, " ")

   c.Groups = groups

   c.Answers = []string{}
	for _, x := range answers {
		c.Answers = append(c.Answers, strings.Join(x, " "))
	}

   c.Notes = []string{}
	for _, x := range notes {
		c.Notes = append(c.Notes, strings.Join(x, " "))
	}

   c.Metadata = algs.New(meta)

   return
}

func (c *Card) HasAnswer() bool {
   return len(c.Answers) > 0
}
