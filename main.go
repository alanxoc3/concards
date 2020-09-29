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
	d := deck.NewDeck(time.Now())

	// We don't care if there is no meta data.
	if predicts, err := file.ReadPredictsFromFile(c.MetaFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error: Unable to open meta file \"%s\".", c.MetaFile)
		os.Exit(1)
	} else {
		d.AddPredicts(predicts...)
	}

	if len(c.Files) == 0 {
		fmt.Fprintf(os.Stderr, "Error: You didn't provide any files to parse.\n")
		os.Exit(1)
	}

	for _, f := range c.Files {
		if cm, err := file.ReadCardsFromFile(f); err != nil {
			fmt.Fprintf(os.Stderr, "Error: File \"%s\" does not exist!\n", f)
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
		lines := d.CardList()

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
	_ = file.WritePredictsToFile(d.PredictList(), c.MetaFile)
}
