package file

import (
   "bufio"
   "fmt"
   "io"
   "os"

   "github.com/alanxoc3/concards/con"
)

// Open opens filename and loads cards into new deck
func ReadToDeck(filename string) (con.Deck, error) {
   if f, err := os.Open(filename); err != nil {
      return nil, fmt.Errorf("Error: Unable to open file \"%s\"", filename)
   } else {
      return ReadToDeckHelper(f)
   }
}

func ReadToDeckHelper(r io.Reader) (d con.Deck, err error) {
   // Initialization.
   d = con.NewDeck()
   facts := [][]string{}
   state := false

   // Scan by words.
   scanner := bufio.NewScanner(r)
   scanner.Split(bufio.ScanWords)

   for scanner.Scan() {
      t := scanner.Text()

      if state {
         if t == "@" {
            facts = append(facts, []string{})
         } else if t == "@>" {
            d = d.AppendCard(facts)
         } else if t == "<@" {
            d = d.AppendCard(facts)
            state = false
         } else {
            if i := len(facts)-1; i >= 0 {
               facts[i] = append(facts[i], t)
            }
         }
      } else if t == "@>" {
         state = true
         facts = [][]string{{}}
      }
   }

   return
}
