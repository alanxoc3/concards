package card

import "testing"
import "fmt"

func TestCard(t *testing.T) {
   c, err := New(
      map[string]bool{"help": true, "bob": true},
      []string{"QUESTION"},
      [][]string{},
      [][]string{},
      []string{})

   if err != nil {
      t.Errorf("Error: %s ", err)
   }

   if c.KeyText() != "@> bob help @q QUESTION" {
      t.Errorf("KeyText generated wrong.")
   }

   correctSum := "9e63f3e18813c7faba4183b74344e6b3e4b47f63fbdfbf633b951da8569d14b7"
   cardSum := fmt.Sprintf("%x", c.Hash())
   if correctSum != cardSum {
      t.Errorf("Hash generated wrong.")
   }

   if c.HasAnswer() {
      t.Errorf("HasAnswer is wrong.")
   }
}
