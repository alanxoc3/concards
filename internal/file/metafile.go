package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/alanxoc3/concards/internal"
	"github.com/alanxoc3/concards/internal/deck"
	"github.com/alanxoc3/concards/internal/meta"
)

// Open opens filename and loads cards into new deck
func ReadMetasToDeck(filename string, d *deck.Deck) error {
	if f, err := os.Open(filename); err != nil {
		return fmt.Errorf("Error: Unable to open meta file \"%s\"", filename)
	} else {
		defer f.Close()
		ReadMetasToDeckHelper(f, d)
		return nil
	}
}

func ReadMetasToDeckHelper(r io.Reader, d *deck.Deck) {
	// Scan by words.
	lineScanner := bufio.NewScanner(r)
	lineScanner.Split(bufio.ScanLines)

	for lineScanner.Scan() {
		strs := strings.Fields(lineScanner.Text())

		// First field is a constant sized checksum.
		if len(strs) > 0 && len(strs[0]) == 32 {
			d.AddMeta(internal.NewHash(strs[0]), meta.NewPredictFromStrings(strs...))
		}
	}
}

func WriteMetasToString(d *deck.Deck) string {
	predicts := d.PredictList()
   predictStrings := make([]string, len(predicts))
   for _, v := range predicts {
      predictStrings = append(predictStrings, v.String())
   }

	sort.Strings(predictStrings)
   return strings.Join(predictStrings, "\n")
}

func WriteMetasToFile(d *deck.Deck, filename string) error {
	str := []byte(WriteMetasToString(d))
	err := ioutil.WriteFile(filename, str, 0644)
	if err != nil {
		return fmt.Errorf("Error: Writing to \"%s\" failed.", filename)
	}

	return nil
}
