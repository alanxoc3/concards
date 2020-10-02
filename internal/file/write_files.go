package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/alanxoc3/concards/internal"
	"github.com/alanxoc3/concards/internal/meta"
)

func createDir(path string) {
	if err := os.MkdirAll(path, 0755); err != nil {
		internal.AssertError(fmt.Sprintf("Problem creating directory \"%s\".", path))
	}
}

func WritePredictsToFile(l []*meta.Predict, filename string) error {
	createDir(filepath.Dir(filename))

	err := ioutil.WriteFile(filename, []byte(WritePredictsToString(l)), 0644)
	if err != nil {
		return fmt.Errorf("Error: Writing to \"%s\" failed.", filename)
	}

	return nil
}

// TODO: clean me up please.
func WriteOutcomesToFile(l []*meta.Outcome, filename string) error {
   outcomes := ReadOutcomesFromFile(filename)
   outcomes = append(outcomes, l...)

   outcomeMap := map[meta.Key]*meta.Outcome{}
   for _, o := range outcomes {
      outcomeMap[o.Key()] = o
   }

	outcomes = []*meta.Outcome{}
	for _, v := range outcomeMap {
		outcomes = append(outcomes, v)
	}

	createDir(filepath.Dir(filename))

	err := ioutil.WriteFile(filename, []byte(WriteOutcomesToString(outcomes)), 0644)
	if err != nil {
		return fmt.Errorf("Error: Writing to \"%s\" failed.", filename)
	}

	return nil
}
