package main

import (
	"os"
	"os/user"

	"github.com/alanxoc3/concards/internal"
	"github.com/spf13/pflag"
)

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

func genConfig(flags *pflag.FlagSet) *internal.Config {
	c := &internal.Config{}
	flags.BoolVarP(&c.IsVersion, "version", "V", false, "Concards build information")
	flags.BoolVarP(&c.IsReview, "review", "r", false, "Show cards available to be reviewed")
	flags.BoolVarP(&c.IsMemorize, "memorize", "m", false, "Show cards available to be memorized")
	flags.BoolVarP(&c.IsDone, "done", "d", false, "Show cards not available to be reviewed or memorized")
	flags.BoolVarP(&c.IsPrint, "print", "p", false, "Print all cards, one card per line")
	flags.BoolVarP(&c.IsFileList, "files-with-cards", "l", false, "Print the file paths that have cards")
	flags.IntVarP(&c.Number, "number", "n", 0, "How many cards to review")
	flags.StringVarP(&c.Editor, "editor", "E", defaultEditor(), "Defaults to \"$EDITOR\" or \"vi\"")
	flags.StringVarP(&c.PredictFile, "predict", "P", "", "Defaults to \"$CONCARDS_PREDICT\" or \"~/.config/concards/predict\"")
	flags.StringVarP(&c.OutcomeFile, "outcome", "O", "", "Defaults to \"$CONCARDS_OUTCOME\" or \"~/.config/concards/outcome\"")
	return c
}

func cleanConfig(c *internal.Config, args []string) {
	if !c.IsReview && !c.IsMemorize && !c.IsDone {
		c.IsReview = true
		c.IsMemorize = true
		c.IsDone = true
	}

	c.Files = args

	if c.Editor      == "" { c.Editor      = defaultEditor() }
	if c.PredictFile == "" { c.PredictFile = defaultEnv("CONCARDS_PREDICT", "predict") }
	if c.OutcomeFile == "" { c.OutcomeFile = defaultEnv("CONCARDS_OUTCOME", "outcome") }
}
