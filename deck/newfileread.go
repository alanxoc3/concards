package deck

import (
	"bufio"
	"fmt"
	"io"
	"os"

)

const (
   OnNothing = iota
   OnExclude = iota
   OnGroup = iota
   OnQuestion = iota
   OnAnswer = iota
   OnNote = iota
   OnTime = iota
   OnMeta = iota
)

type parseinfo struct {
   wordsBuffer []string
   excludes map[string]bool
   groups map[string]bool
   question []string
   answers [][]string
   notes [][]string
   meta []string
   timestamp []string
   state int
   prevState int
}

func NewInfo() *parseinfo {
   return &parseinfo{
      excludes: map[string]bool{},
      groups: map[string]bool{},
      question: []string{},
      answers: [][]string{},
      notes: [][]string{},
      timestamp: []string{},
      meta: []string{},
      state: OnNothing,
      prevState: OnNothing }
}

var ParseValues = map[string]int{
   "@>": OnGroup,
   "<@": OnNothing,
   "@!": OnExclude,
   "@q": OnQuestion, "@Q": OnQuestion,
   "@a": OnAnswer,   "@A": OnAnswer,
   "@n": OnNote,     "@N": OnNote,
   "@t": OnTime,     "@T": OnTime,
   "@m": OnMeta,     "@M": OnMeta }

func UpdatePrevState(info *parseinfo) {
   switch info.state {
   case OnNothing, OnGroup, OnQuestion:
      info.prevState = info.state
   }
}

// Open opens filename and loads cards into new deck
func OpenNewFormat(filename string) (d *DeckControl, err error) {
	d = &DeckControl{}

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

   info := NewInfo()

   for scanner.Scan() {
      t := scanner.Text()

      if info.state != OnNothing || info.state == OnNothing && ParseValues[t] == OnGroup {
         if val, ok := ParseValues[t]; ok {
            UpdatePrevState(info)
            info.state = val

            switch val {
               case OnGroup: greaterthan_logic(info)
               case OnNothing:
                  lessthan_logic(info)
               case OnExclude: exclude_logic(info)
               case OnQuestion: question_logic(info)
               case OnAnswer: answer_logic(info)
               case OnNote: note_logic(info)
               case OnTime: timestamp_logic(info)
               case OnMeta: meta_logic(info)
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
                  i := len(info.answers)-1
                  info.answers[i] = append(info.answers[i], t)
               } else if info.state == OnNote {
                  i := len(info.notes)-1
                  info.notes[i] = append(info.notes[i], t)
               } else if info.state == OnTime {
                  info.timestamp = append(info.timestamp, t)
               } else if info.state == OnMeta {
                  info.meta = append(info.meta, t)
               }
            }
         }
      }
   }

   return
}

func greaterthan_logic(info *parseinfo) {
   switch info.prevState {
      case OnNothing: info = NewInfo() // reset the info
      case OnGroup: break // continue adding to the group.
      case OnQuestion:
         // TODO: add the question, then reset it
         info.question = []string{}
         info.answers = [][]string{}
         info.notes = [][]string{}
         info.timestamp = []string{}
         info.meta = []string{}
   }
}

func lessthan_logic(info *parseinfo) {
   switch info.prevState {
      case OnNothing: break
      case OnGroup: break
      case OnQuestion:
         // TODO: add the question, then reset it
         info.question = []string{}
         info.answers = [][]string{}
         info.notes = [][]string{}
         info.timestamp = []string{}
         info.meta = []string{}
   }
}

func exclude_logic(info *parseinfo) {
   switch info.prevState {
      case OnNothing: break
      case OnGroup: fallthrough
      case OnQuestion:
         info.excludes = map[string]bool{}
   }
}

func question_logic(info *parseinfo) {
   switch info.prevState {
      case OnNothing: break
      case OnGroup: break
      case OnQuestion:
         // TODO: add question, then reset it
         info.question = []string{}
         info.answers = [][]string{}
         info.notes = [][]string{}
         info.timestamp = []string{}
         info.meta = []string{}
   }
}

func answer_logic(info *parseinfo) {
   switch info.prevState {
      case OnNothing: break
      case OnGroup: break // do nothing for answers following groups (for now)
      case OnQuestion:
         info.answers = append(info.answers, []string{})
   }
}

func note_logic(info *parseinfo) {
   switch info.prevState {
      case OnNothing: break
      case OnGroup: break // TODO: notes for group in future?
      case OnQuestion:
         info.notes = append(info.notes, []string{})
   }
}

func timestamp_logic(info *parseinfo) {
   switch info.prevState {
      case OnNothing: break
      case OnGroup: break // TODO: meta for group in future?
      case OnQuestion:
         info.timestamp = []string{}
   }
}

func meta_logic(info *parseinfo) {
   switch info.prevState {
      case OnNothing: break
      case OnGroup: break // TODO: meta for group in future?
      case OnQuestion:
         info.meta = []string{}
   }
}


// There is a forward looking approach, and there is a backwards approach.
// Group and exclude are both already doing the forward approach. I like it
// too. I'll try it. I think I did the backwards approach for the previous one.
