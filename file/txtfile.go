package file

import (
   "bufio"
   "fmt"
   "io"
   "os"
   "path/filepath"

   "github.com/alanxoc3/concards/core"
)

// Open opens filename and loads cards into new deck
func ReadCardsToDeck(filename string, d *core.Deck) error {
   filename, _ = filepath.Abs(filename)
   if f, fe := os.Open(filename); fe != nil {
      return fmt.Errorf("Error: Unable to open file \"%s\"", filename)
   } else if s, se := os.Lstat(filename); se != nil {
      return fmt.Errorf("Error: Unable to open file \"%s\"", filename)
   } else if s.IsDir() {
      return fmt.Errorf("Error: Unable to open folder \"%s\"", filename)
   } else {
      ReadCardsToDeckHelper(f, d, filename)
      return nil
   }
}

func ReadCardsToDeckHelper(r io.Reader, d *core.Deck, f string) {
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
            d.AddFacts(facts, f)
            facts = [][]string{{}}
         } else if t == "<@" {
            d.AddFacts(facts, f)
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
