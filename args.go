// "pre-xxxx" hooks in git do affect the process.
// "post-xxxx" hooks in git don't affect the process.
package main

import (
	"os"
	"os/user"
        "path/filepath"

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

func getDataDir() string {
	if val, present := os.LookupEnv("CONCARDS_DATA_DIR"); present {
		return val
	} else if val, present := os.LookupEnv("XDG_DATA_HOME"); present {
		return val
	} else if usr, err := user.Current(); err == nil {
		return usr.HomeDir + "/.local/share/concards"
	} else {
		return ""
	}
}

func getConfigDir() string {
	if val, present := os.LookupEnv("CONCARDS_CONFIG_DIR"); present {
		return val
	} else if val, present := os.LookupEnv("XDG_CONFIG_HOME"); present {
		return filepath.Join(val, "concards")
	} else if usr, err := user.Current(); err == nil {
		return usr.HomeDir + "/.config/concards"
	} else {
		return ""
	}
}

func genConfig(flags *pflag.FlagSet) *internal.Config {
	c := &internal.Config{}
	flags.BoolVarP(&c.IsVersion, "version", "V", false, "Concards build information.")
	flags.BoolVarP(&c.IsReview, "review", "r", false, "Show cards available to be reviewed.")
	flags.BoolVarP(&c.IsMemorize, "memorize", "m", false, "Show cards available to be memorized.")
	flags.BoolVarP(&c.IsDone, "done", "d", false, "Show cards not available to be reviewed or memorized.")
	flags.BoolVarP(&c.IsPrint, "print", "p", false, "Print all cards, one card per line.")
	flags.BoolVarP(&c.IsFileList, "files-with-cards", "l", false, "Print the file paths that have cards.")
	flags.IntVarP(&c.Number, "number", "n", 0, "How many cards to review.")
	flags.StringVarP(&c.Editor, "editor", "E", defaultEditor(), "Defaults to \"$EDITOR\" or \"vi\".")
	flags.StringVar(&c.DataDir, "data-dir", "", "Override the data directory location.")
	flags.StringVar(&c.ConfigDir, "config-dir", "", "Override the config directory location.")
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
	if c.DataDir     == "" { c.DataDir = getDataDir() }
	if c.ConfigDir   == "" { c.ConfigDir = getConfigDir() }

	c.PredictFile = filepath.Join(c.DataDir, "predict")
	c.OutcomeFile = filepath.Join(c.DataDir, "outcome")

	c.EventReviewFile = filepath.Join(c.ConfigDir, "hooks", "event-review")
	c.EventStartupFile = filepath.Join(c.ConfigDir, "hooks", "event-startup")
}
