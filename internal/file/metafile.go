package file

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/alanxoc3/concards/internal/meta"
)

// TODO: Clean up duplicate code here.

// Open filename and loads cards into new deck
func ReadPredictsFromFile(filename string) []*meta.Predict {
	f, err := os.Open(filename)
	if err != nil {
		return []*meta.Predict{}
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

// Open filename and loads cards into new deck
func ReadOutcomesFromFile(filename string) []*meta.Outcome {
	f, err := os.Open(filename)
	if err != nil {
		return []*meta.Outcome{}
	}

	defer f.Close()
	return ReadOutcomesFromReader(f)
}

func ReadOutcomesFromReader(r io.Reader) []*meta.Outcome {
	// Scan by words.
	lineScanner := bufio.NewScanner(r)
	lineScanner.Split(bufio.ScanLines)
	list := []*meta.Outcome{}

	for lineScanner.Scan() {
		strs := strings.Fields(lineScanner.Text())

		// Only add if there is something on the line.
		if len(strs) > 0 {
			list = append(list, meta.NewOutcomeFromStrings(strs...))
		}
	}

	return list
}

func WriteOutcomesToString(l []*meta.Outcome) string {
	outcomeStrings := []string{}
	for _, v := range l {
		if !v.IsZero() {
			outcomeStrings = append(outcomeStrings, v.String())
		}
	}

	sort.Strings(outcomeStrings)
	return strings.Join(outcomeStrings, "\n")
}
