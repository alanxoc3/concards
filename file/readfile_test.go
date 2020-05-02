package file

import "testing"
import "strings"

func TestCard(t *testing.T) {
   r := strings.NewReader(" @> hi @ hello <@")
   d, err := ReadToDeckHelper(r)
   if err != nil {
      println(err.Error())
   }

	for _, c := range d {
      println("card!")
      for _, f := range c {
         println("- " + f)
      }
	}
}
