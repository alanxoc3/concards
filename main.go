package main

import (
	"fmt"
	"os"
	"os/user"
	// "math/rand"
	// "time"
	"strings"

	"github.com/spf13/cobra"
	// "github.com/alanxoc3/concards/internal"
	// "github.com/alanxoc3/concards/internal/deck"
	// "github.com/alanxoc3/concards/internal/file"
	// "github.com/alanxoc3/concards/internal/termboxgui"
)

var version string = "snapshot"

type Config struct {
	// The various true or false options
	IsReview   bool
	IsMemorize bool
	IsDone     bool
	IsPrint    bool
	IsFileList bool
	IsStream   bool

	Editor      string
	Number      int
	PredictFile string
	OutcomeFile string
	Files       []string
}

func defaultEditor() string {
	if val, present := os.LookupEnv("EDITOR"); present {
		return val
	} else {
		return "vi"
	}
}

func defaultEnv(env string, file string) string {
	if val, present := os.LookupEnv(env); present {
		return val
	} else if usr, err := user.Current(); err == nil {
		return usr.HomeDir + "/.config/concards/" + file
	} else {
		return ""
	}
}

func main() {
	//    config := Config{}
	var version bool
	var review bool
	var memorize bool
	var done bool
	var print bool
	var filelist bool
	var number int
	var editor string
	var predictfile string
	var outcomefile string

	var rootCmd = &cobra.Command{Use: "concards"}

	rootCmd.Flags().BoolVarP(&version, "version", "V", false, "Concards build information")
	rootCmd.Flags().BoolVarP(&review, "review", "r", false, "Show cards available to be reviewed")
	rootCmd.Flags().BoolVarP(&memorize, "memorize", "m", false, "Show cards available to be memorized")
	rootCmd.Flags().BoolVarP(&done, "done", "d", false, "Show cards not available to be reviewed or memorized")
	rootCmd.Flags().BoolVarP(&print, "print", "p", false, "Print all cards, one card per line")
	rootCmd.Flags().BoolVarP(&filelist, "filelist", "l", false, "Print the file paths that have cards")
	rootCmd.Flags().IntVarP(&number, "number", "n", 0, "How many cards to review")
	rootCmd.Flags().StringVarP(&editor, "editor", "E", defaultEditor(), "Defaults to \"$EDITOR\" or \"vi\"")
	rootCmd.Flags().StringVarP(&predictfile, "predict", "P", defaultEnv("CONCARDS_PREDICT", "predict"), "Defaults to \"$CONCARDS_PREDICT\" or \"~/.config/concards/predict\"")
	rootCmd.Flags().StringVarP(&outcomefile, "outcome", "O", defaultEnv("CONCARDS_OUTCOME", "outcome"), "Defaults to \"$CONCARDS_OUTCOME\" or \"~/.config/concards/outcome\"")

	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		fmt.Println("hello: " + strings.Join(args, " "))
	}
	// rootCmd.AddCommand(cmdPrint, cmdEcho)
	// cmdEcho.AddCommand(cmdTimes)
	rootCmd.Execute()
}

/*
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
*/
