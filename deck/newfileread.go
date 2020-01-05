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

// Search through the file. If there is an \n@>\s 

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

   curState := OnNothing

   exSet := make(map[string]bool)
   groups := make(map[string]bool)

	for scanner.Scan() {
      t := scanner.Text()

      if curState == OnNothing {
         if t == "@>" {
            curState = OnGroup
            groups = make(map[string]bool)
         }
      } else {
         switch t {
         case "@>":
            curState = OnGroup
            groups = make(map[string]bool)
         case "<@":
            curState = OnNothing
            fmt.Println()
         case "@!":
            exSet = make(map[string]bool)
            curState = OnExclude
         case "@q", "@Q": curState = OnQuestion
         case "@a", "@A": curState = OnAnswer
         case "@t", "@T": curState = OnTime
         case "@n", "@N": curState = OnNote
         case "@m", "@M": curState = OnMeta
         default:
            if exSet[t] {
            } else if curState == OnGroup {
               groups[t] = true
            } else if curState == OnExclude {
               exSet[t] = true
            } else if curState == OnQuestion {
               fmt.Println(t)
            } else if curState == OnAnswer {
               fmt.Println("\t" + t)
            }
         }
      }
   }

	return
}
