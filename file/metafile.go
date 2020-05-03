package file

import (
   "bufio"
   "strings"
   "io/ioutil"
   "sort"
   "fmt"
   "io"
   "os"

   "github.com/alanxoc3/concards/core"
)

func oe(arr []string, i int) string {
   if i < len(arr) {
      return arr[i]
   } else {
      return ""
   }
}

// Open opens filename and loads cards into new deck
func ReadMetasToDeck(filename string, d *core.Deck) (*core.Deck, error) {
   if f, err := os.Open(filename); err != nil {
      return nil, fmt.Errorf("Error: Unable to open file \"%s\"", filename)
   } else {
      return ReadMetasToDeckHelper(f, d), nil
   }
}

func ReadMetasToDeckHelper(r io.Reader, d *core.Deck) *core.Deck {
   // Scan by words.
   line_scanner := bufio.NewScanner(r)
   line_scanner.Split(bufio.ScanLines)

   for line_scanner.Scan() {
      strs := strings.Fields(line_scanner.Text())

      // First field is a constant sized checksum.
      if len(strs) > 0 && len(strs[0]) == 32 {
         d.AddMeta(strs[0], core.NewMeta(oe(strs, 1), oe(strs, 2), oe(strs, 3), strs[4:]))
      }
   }

   return d
}

func WriteMetasToString(d *core.Deck) (file_str string) {
   // Copy keys
   keys := make([]string, len(d.Mmap))

   i := 0
   for k := range d.Mmap {
       keys[i] = k
       i++
   }

   // Sort keys
   sort.Strings(keys)

   // Create string
   for _, k := range keys {
      m := d.Mmap[k]
      if m != nil && m.String() != "" {
         file_str += m.String() + "\n"
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
