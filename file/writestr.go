package file

import (
	"fmt"
	"strings"
	"io/ioutil"

	"github.com/alanxoc3/concards/deck"
	"github.com/alanxoc3/concards/card"
)

func WriteCardToString(c *card.Card) (str string) {
   str = "@> " + strings.Join(c.Groups.ToArray(), " ")
	str += "\n@q " + c.Question

   for _, a := range c.Answers {
		str += "\n@a " + a
   }

   str += "\n@m " + c.Metadata.ToString()

	return str
}

func WriteDeckToString(d *deck.Deck) (str string) {
	// do groups stuff
	for _, c := range *d {
		str += WriteCardToString(c) + "\n\n"
	}

	return
}

func WriteDeckToFile(d *deck.Deck, filename string, message string) error {
	str := []byte(message + "\n\n" + WriteDeckToString(d))
	err := ioutil.WriteFile(filename, str, 0644)
	if err != nil {
		return fmt.Errorf("Error: Writing to \"%s\" failed.", filename)
	}

	return nil
}
