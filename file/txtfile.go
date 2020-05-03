package file

import (
   "bufio"
   "fmt"
   "io"
   "os"

   "github.com/alanxoc3/concards/core"
)

// Open opens filename and loads cards into new deck
func ReadCardsToDeck(filename string, d *core.Deck) error {
   if f, err := os.Open(filename); err != nil {
      return fmt.Errorf("Error: Unable to open file \"%s\"", filename)
   } else {
      ReadCardsToDeckHelper(f, d)
      return nil
   }
}

func ReadCardsToDeckHelper(r io.Reader, d *core.Deck) {
   // Initialization.
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
            d.AddFacts(facts)
            facts = [][]string{{}}
         } else if t == "<@" {
            d.AddFacts(facts)
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
