package file

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/alanxoc3/concards/deck"
	"github.com/alanxoc3/concards/card"
)

const (
   OnNothing = iota
   OnExclude = iota
   OnGroup = iota
   OnQuestion = iota
   OnAnswer = iota
   OnNote = iota
   OnMeta = iota
)

type parseinfo struct {
   excludes map[string]bool
   groups map[string]bool
   question []string
   answers [][]string
   notes [][]string
   meta []string
   state int
   prevState int
   shouldCreateCard bool
}

func newInfo() *parseinfo {
   return &parseinfo{
      excludes: map[string]bool{},
      groups: map[string]bool{},
      question: []string{},
      answers: [][]string{},
      notes: [][]string{},
      meta: []string{},
      state: OnNothing,
      prevState: OnNothing }
}

func resetGroup(info *parseinfo) {
   info.groups = map[string]bool{}
}

func resetCard(info *parseinfo) {
   info.question = []string{}
   info.answers = [][]string{}
   info.notes = [][]string{}
   info.meta = []string{}
   info.shouldCreateCard = false
}

// Returns true if we have a Byte Order Marker at the beginning of the file.
// TODO: Don't make this just a hotfix, get rid of any marker, not just UTF-8.
func isBOM(bom []byte) bool {
	return bom[0] == 0xef && bom[1] == 0xbb && bom[2] == 0xbf
}

var ParseValues = map[string]int{
   "@>": OnGroup,
   "<@": OnNothing,
   "@!": OnExclude,
   "@q": OnQuestion, "@Q": OnQuestion,
   "@a": OnAnswer,   "@A": OnAnswer,
   "@n": OnNote,     "@N": OnNote,
   "@m": OnMeta,     "@M": OnMeta }

func UpdatePrevState(info *parseinfo) {
   switch info.state {
   case OnNothing, OnGroup, OnQuestion:
      info.prevState = info.state
   }
}

// Open opens filename and loads cards into new deck
func ReadToDeck(filename string) (d deck.Deck, err error) {
	d = deck.Deck{}

	file, err1 := os.Open(filename)
	if err1 != nil {
		err = fmt.Errorf("Error: Unable to open file \"%s\"", filename)
		return
	}

	// Get rid of UTF-8 encoding.
	bom := make([]byte, 3, 3)

	// Returns an error if fewer than 3 bytes were read.
	io.ReadFull(file, bom[:])
	if !isBOM(bom) {
		file.Seek(0, 0)
	}

	// Set up the line stuff
	scanner := bufio.NewScanner(file)
   scanner.Split(bufio.ScanWords)

   info := newInfo()

   for scanner.Scan() {
      t := scanner.Text()

      if info.state != OnNothing || info.state == OnNothing && ParseValues[t] == OnGroup {
         if val, ok := ParseValues[t]; ok {
            UpdatePrevState(info)
            info.state = val

            switch val {
               case OnGroup: greaterthan_logic(info)
               case OnNothing: lessthan_logic(info)
               case OnExclude: exclude_logic(info)
               case OnQuestion: question_logic(info)
               case OnAnswer: answer_logic(info)
               case OnNote: note_logic(info)
               case OnMeta: meta_logic(info)
            }

            if info.shouldCreateCard {
               if c, err := card.New(info.groups, info.question, info.answers, info.notes, info.meta); err == nil {
                  d = append(d, c)
               }
               resetCard(info)
            }
         } else {
            if info.state != OnNothing {
               if info.excludes[t] {
               } else if info.state == OnExclude {
                  info.excludes[t] = true
               } else if info.state == OnGroup {
                  info.groups[t] = true
               } else if info.state == OnQuestion {
                  info.question = append(info.question, t)
               } else if info.state == OnAnswer {
                  if i := len(info.answers)-1; i > 0 {
                     info.answers[i] = append(info.answers[i], t)
                  }
               } else if info.state == OnNote {
                  if i := len(info.notes)-1; i > 0 {
                     info.notes[i] = append(info.notes[i], t)
                  }
               } else if info.state == OnMeta {
                  info.meta = append(info.meta, t)
               }
            }
         }
      }
   }

   return
}
