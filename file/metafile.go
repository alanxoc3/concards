package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/alanxoc3/concards/core"
)

func getParam(arr []string, i int) string {
	if i < len(arr) {
		return arr[i]
	} else {
		return ""
	}
}

// Open opens filename and loads cards into new deck
func ReadMetasToDeck(filename string, d *core.Deck) error {
	if f, err := os.Open(filename); err != nil {
		return fmt.Errorf("Error: Unable to open meta file \"%s\"", filename)
	} else {
		defer f.Close()
		ReadMetasToDeckHelper(f, d)
		return nil
	}
}

func ReadMetasToDeckHelper(r io.Reader, d *core.Deck) {
	// Scan by words.
	lineScanner := bufio.NewScanner(r)
	lineScanner.Split(bufio.ScanLines)

	for lineScanner.Scan() {
		strs := strings.Fields(lineScanner.Text())

		// First field is a constant sized checksum.
		if len(strs) > 0 && len(strs[0]) == 32 {
			d.AddMeta(strs[0], core.NewMeta(getParam(strs, 1), getParam(strs, 2), getParam(strs, 3), strs[4:]))
		}
	}
}

func WriteMetasToString(d *core.Deck) (fileStr string) {
	// Copy keys
	keys := make([]string, len(d.MetaMap))

	i := 0
	for k := range d.MetaMap {
		keys[i] = k
		i++
	}

	// Sort keys
	sort.Strings(keys)

	// Create string
	for _, k := range keys {
		m := d.MetaMap[k]
		if m != nil && m.String() != "" {
			fileStr += k + " " + m.String() + "\n"
		}
	}

	return
}

func WriteMetasToFile(d *core.Deck, filename string) error {
	str := []byte(WriteMetasToString(d))
	err := ioutil.WriteFile(filename, str, 0644)
	if err != nil {
		return fmt.Errorf("Error: Writing to \"%s\" failed.", filename)
	}

	return nil
}
