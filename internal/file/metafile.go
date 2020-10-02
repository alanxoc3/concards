package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/alanxoc3/concards/internal"
	"github.com/alanxoc3/concards/internal/meta"
)

// Open filename and loads cards into new deck
func ReadPredictsFromFile(filename string) []*meta.Predict {
	f, err := os.Open(filename)
	if err != nil {
		WritePredictsToFile([]*meta.Predict{}, filename)
		f, err = os.Open(filename)
		if err != nil {
         fmt.Println(err)
			internal.AssertError(fmt.Sprintf("Couldn't access file \"%s\".", filename))
		}
	}

	defer f.Close()
	return ReadPredictsFromReader(f)
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

func WritePredictsToString(l []*meta.Predict) string {
	predictStrings := []string{}
	for _, v := range l {
		if !v.IsZero() {
			predictStrings = append(predictStrings, v.String())
		}
	}

	sort.Strings(predictStrings)
	return strings.Join(predictStrings, "\n")
}
