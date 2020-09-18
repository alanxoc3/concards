package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"

	"github.com/alanxoc3/concards/internal/deck"
	"github.com/alanxoc3/concards/internal/file"
	"github.com/alanxoc3/concards/internal/termboxgui"
)

var version string = "snapshot"

func main() {
	c := file.GenConfig(version)
	d := deck.NewDeck(time.Now())

	// We don't care if there is no meta data.
	file.ReadMetasToDeck(c.MetaFile, d)

	if len(c.Files) == 0 {
		fmt.Printf("Error: You didn't provide any files to parse.\n")
		os.Exit(1)
	}

	for _, f := range c.Files {
		if cm, err := file.ReadCards(f); err != nil {
			fmt.Printf("Error: File \"%s\" does not exist!\n", f)
			os.Exit(1)
		} else {
			for _, c := range cm {
				d.AddCards(c)
			}
		}
	}

	if !c.IsMemorize {
		d.RemoveMemorize()
	}
	if !c.IsReview {
		d.RemoveReview()
	}
	if !c.IsDone {
		d.RemoveDone()
	}
	if c.Number > 0 {
		d.Truncate(c.Number)
	}

	if c.IsPrint {
		lines := []string{}
		for _, v := range d.CardList() {
			lines = append(lines, v.String())
		}

		sort.Strings(lines)

		for _, c := range lines {
			fmt.Printf("@> %s\n", c)
		}
		if len(lines) > 0 {
			fmt.Printf("<@\n")
		}
		return
	}

	rand.Seed(time.Now().UTC().UnixNano()) // Used for algorithms.
	termboxgui.TermBoxRun(d, c)
	_ = file.WriteMetasToFile(d, c.MetaFile)
}
