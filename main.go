package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/alanxoc3/concards/internal/deck"
	"github.com/alanxoc3/concards/internal/file"
	"github.com/alanxoc3/concards/internal/termboxgui"
)

var version string = "snapshot"

func main() {
	c := file.GenConfig(version)
	d := deck.NewDeck()

	// We don't care if there is no meta data.
	file.ReadMetasToDeck(c.MetaFile, d)

	if len(c.Files) == 0 {
		fmt.Printf("Error: You didn't provide any files to parse.\n")
		os.Exit(1)
	}

	for _, f := range c.Files {
		if err := file.ReadCardsToDeck(d, f); err != nil {
			fmt.Printf("Error: File \"%s\" does not exist!\n", f)
			os.Exit(1)
		}
	}

	if !c.IsMemorize {
		d.FilterOutMemorize()
	}
	if !c.IsReview {
		d.FilterOutReview()
	}
	if !c.IsDone {
		d.FilterOutDone()
	}
	if c.Number > 0 {
		d.FilterNumber(c.Number)
	}

	if c.IsPrint {
		for i := 0; i < d.Len(); i++ {
			fmt.Printf("@> %s\n", d.GetCard(i).String())
		}

		if d.Len() > 0 {
			fmt.Printf("<@\n")
		}
		return
	}

	rand.Seed(time.Now().UTC().UnixNano())
	d.Shuffle()
	termboxgui.TermBoxRun(d, c)
	_ = file.WriteMetasToFile(d, c.MetaFile)
}
