package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/alanxoc3/concards/internal"
	"github.com/alanxoc3/concards/internal/deck"
	"github.com/alanxoc3/concards/internal/file"
	"github.com/alanxoc3/concards/internal/termboxgui"
	"github.com/spf13/cobra"
)

var version string = "snapshot"

func main() {
	rootCmd := &cobra.Command{
		Use:   "concards [flags] [file | folder]...",
		Short: "Concards is a simple CLI based SRS flashcard app.",
	}

	c := genConfig(rootCmd.Flags())
	rootCmd.PreRun = func(cmd *cobra.Command, args []string) {
		cleanConfig(c, args)
	}

	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		if c.IsVersion {
			fmt.Println("concards " + version)
			return
		}
		main_logic(c)
	}

	rootCmd.Execute()
}

func main_logic(c *internal.Config) {
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
