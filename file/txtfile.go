package file

import (
   "bufio"
   "strings"
   "io"
   "fmt"
   "os"
   "path/filepath"

   "github.com/alanxoc3/concards/core"
)

// Open opens filename and loads cards into new deck
func ReadCardsToDeck(d *core.Deck, filename string) error {
   err := filepath.Walk(filename, func(path string, info os.FileInfo, e error) error {
      if e != nil {
         return e
      }

      n := info.Name()
      is_hidden := len(n) > 1 && string(n[0]) == "." && n != ".."
      is_dir := info.IsDir()

      if is_dir && is_hidden {
         return filepath.SkipDir
      } else if is_hidden || is_dir {
         return nil
      }

      abs_path, _ := filepath.Abs(path)
      if f, fe := os.Open(abs_path); fe != nil {
         return fmt.Errorf("Error: Unable to open file \"%s\"", filename)
      } else {
         defer f.Close()
         ReadCardsToDeckHelper(f, d, abs_path)
      }

      return nil
   })

   return err
}

func ReadCardsToDeckHelper(r io.Reader, d *core.Deck, f string) {
   // Initialization.
   facts := []string{}
   state := false
   var td *core.Deck

   // Scan by words.
   scanner := bufio.NewScanner(r)
   scanner.Split(bufio.ScanWords)

   for scanner.Scan() {
      t := scanner.Text()

      if state {
         if t == "@>" {
            td.AddCardAndSubCardsFromSides(f, strings.Join(facts, " "))
            facts = []string{}
         } else if t == "<@" {
            td.AddCardAndSubCardsFromSides(f, strings.Join(facts, " "))
            for i := 0; i < td.Len(); i++ {
               d.AddCard(td.GetCard(i))
            }
            state = false
         } else {
            facts = append(facts, t)
         }
      } else if t == "@>" {
         // create td
         td = core.NewDeck()
         state = true
         facts = []string{}
      }
   }

   return
}
