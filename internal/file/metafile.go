package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/alanxoc3/concards/internal/deck"
	"github.com/alanxoc3/concards/internal/meta"
)

// Open filename and loads cards into new deck
func ReadPredictsFromFile(filename string) ([]*meta.Predict, error) {
	if f, err := os.Open(filename); err != nil {
		return nil, fmt.Errorf("Error: Unable to open meta file \"%s\".", filename)
	} else {
		defer f.Close()
		return ReadPredictsFromReader(f), nil
	}
}

func ReadPredictsFromReader(r io.Reader) []*meta.Predict {
	// Scan by words.
	lineScanner := bufio.NewScanner(r)
	lineScanner.Split(bufio.ScanLines)
   list := []*meta.Predict{}

	for lineScanner.Scan() {
		strs := strings.Fields(lineScanner.Text())

		// Only add if there is something on the line.
		if len(strs) > 0 {
         list = append(list, meta.NewPredictFromStrings(strs...))
		}
	}

   return list
}

func WritePredictsToFile(d *deck.Deck, filename string) error {
	str := []byte(WritePredictsToString(d))
	err := ioutil.WriteFile(filename, str, 0644)
	if err != nil {
		return fmt.Errorf("Error: Writing to \"%s\" failed.", filename)
	}

	return nil
}

func WritePredictsToString(d *deck.Deck) string {
	predicts := d.PredictList()
	predictStrings := make([]string, len(predicts))
	for _, v := range predicts {
		predictStrings = append(predictStrings, v.String())
	}

	sort.Strings(predictStrings)
	return strings.Join(predictStrings, "\n")
}
