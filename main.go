package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/alanxoc3/concards/internal"
	"github.com/alanxoc3/concards/internal/deck"
	"github.com/alanxoc3/concards/internal/file"
	"github.com/alanxoc3/concards/internal/termboxgui"
)

var version string = "snapshot"

func main() {
	c := file.GenConfig(version)
	d := deck.NewDeck(time.Now())

	predicts := file.ReadPredictsFromFile(c.PredictFile)
	d.AddPredicts(predicts...)

	if len(c.Files) == 0 {
		internal.AssertError("You didn't provide any files to parse.")
	}

	for _, f := range c.Files {
		if cm, err := file.ReadCardsFromFile(f); err != nil {
			internal.AssertError(fmt.Sprintf("File \"%s\" does not exist!", f))
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

	if c.IsFileList {
		files := map[string]bool{}

		for _, c := range d.CardList() {
			file := c.File()
			if _, exist := files[file]; !exist {
				files[file] = true
				fmt.Printf("%s\n", file)
			}
		}

		return
	} else if c.IsPrint {
		cards := d.CardList()

		for _, c := range cards {
			fmt.Printf("#: %s\n", c)
		}

		if len(cards) > 0 {
			fmt.Printf(":#\n")
		}

		return
	}

	rand.Seed(time.Now().UTC().UnixNano()) // Used for algorithms.
	termboxgui.TermBoxRun(d, c)
	_ = file.WritePredictsToFile(d.PredictList(), c.PredictFile)
	_ = file.WriteOutcomesToFile(d.OutcomeList(), c.OutcomeFile)
}
