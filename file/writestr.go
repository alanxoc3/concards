package file

// import (
// 	"fmt"
// 	"strings"
// 	"time"
// 	"io/ioutil"
// 
// 	"github.com/alanxoc3/concards/deck"
// 	"github.com/alanxoc3/concards/card"
// )
// 
// func WriteCardToString(c *card.Card) (str string) {
//    str = "@> " + strings.Join(c.Groups.ToArray(), " ")
// 	str += "\n@q " + c.Question
// 
//    for _, x := range c.Answers {
// 		str += "\n@a " + x
//    }
// 
//    for _, x := range c.Notes {
// 		str += "\n@n " + x
//    }
// 
//    str += fmt.Sprintf("\n@m %s %s %d", c.Metadata.Name, c.Metadata.Next.Format(time.RFC3339), c.Metadata.Streak)
//    for _, v := range c.Metadata.Params {
//       str += " " + v
//    }
// 
// 	return str
// }
// 
// func WriteDeckToString(d *deck.Deck) (str string) {
// 	// do groups stuff
// 	for _, c := range *d {
// 		str += WriteCardToString(c) + "\n\n"
// 	}
// 
// 	return
// }
// 
// func WriteDeckToFile(d *deck.Deck, filename string, message string) error {
// 	str := []byte(message + "\n\n" + WriteDeckToString(d))
// 	err := ioutil.WriteFile(filename, str, 0644)
// 	if err != nil {
// 		return fmt.Errorf("Error: Writing to \"%s\" failed.", filename)
// 	}
// 
// 	return nil
// }
