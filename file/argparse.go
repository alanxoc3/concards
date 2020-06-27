package file

import (
	"fmt"
	"github.com/alanxoc3/argparse"
	"os"
	"os/user"
)

type Config struct {
	// The various true or false options
	IsReview   bool
	IsMemorize bool
	IsDone     bool
	IsSides    bool
	IsPrint    bool
	IsStream   bool

	Editor   string
	Number   int
	MetaFile string
	Files    []string
}

// For debugging
func (c *Config) String() string {
	return fmt.Sprintf(`IsReview   %t
IsMemorize %t
IsDone     %t
IsPrint    %t
IsStream   %t
Editor     "%s"
Number     %d
MetaFile   "%s"
Files      %s`, c.IsReview, c.IsMemorize, c.IsDone, c.IsPrint, c.IsStream, c.Editor, c.Number, c.MetaFile, c.Files)
}

func getDefaultEditor() string {
	if val, present := os.LookupEnv("EDITOR"); present {
		return val
	} else {
		return "vi"
	}
}

func getDefaultMeta() string {
	if val, present := os.LookupEnv("CONCARDS_META"); present {
		return val
	} else if usr, err := user.Current(); err == nil {
		return usr.HomeDir + "/.concards-meta"
	} else {
		return ".concards-meta"
	}
}

func GenConfig() *Config {
	// Create new parser object
	parser := argparse.NewParser("concards", "Concards is a simple CLI based SRS flashcard app.")

	// Create flags
	f_version := parser.Flag("V", "version", &argparse.Options{Help: "Concards build information."})
	f_review := parser.Flag("r", "review", &argparse.Options{Help: "Show cards available to be reviewed."})
	f_memorize := parser.Flag("m", "memorize", &argparse.Options{Help: "Show cards available to be memorized."})
	f_done := parser.Flag("d", "done", &argparse.Options{Help: "Show cards not available to be reviewed or memorized."})
	f_sides := parser.Flag("s", "sides", &argparse.Options{Help: "Add cards for all sides."})
	f_print := parser.Flag("p", "print", &argparse.Options{Help: "Prints all cards, one line per card."})
	f_number := parser.Int("n", "number", &argparse.Options{Default: 0, Help: "How many cards to review."})
	f_editor := parser.String("E", "editor", &argparse.Options{Default: getDefaultEditor(), Help: "Which editor to use. Defaults to \"$EDITOR\""})
	f_meta := parser.String("M", "meta", &argparse.Options{Default: getDefaultMeta(), Help: "Path to meta file. Defaults to \"$CONCARDS_META\" or \"~/.concards-meta\""})

	parser.HelpFunc = func(c *argparse.Command, msg interface{}) string {
		var help string
		help += fmt.Sprintf("%s\n\nUsage:\n  %s [OPTIONS]... [FILE|FOLDER]...\n\nOptions:\n", c.GetDescription(), c.GetName())

		for _, arg := range c.GetArgs() {
			if arg.IsFlag() {
				help += fmt.Sprintf("  -%s  --%-9s %s.\n", arg.GetSname(), arg.GetLname(), arg.GetOpts().Help)
			} else {
				help += fmt.Sprintf("  -%s  --%-9s %s.\n", arg.GetSname(), arg.GetLname()+" "+arg.GetSname(), arg.GetOpts().Help)
			}
		}

		return help
	}

	// Parse input
	files, err := parser.ParseReturnArguments(os.Args)
	if err != nil {
		fmt.Print(parser.Help(nil))
		os.Exit(1)
	}

	if *f_version {
		fmt.Printf("Concards v2.0\n")
		os.Exit(0)
	}

	c := &Config{}

	c.IsReview = *f_review
	c.IsMemorize = *f_memorize
	c.IsDone = *f_done
	c.IsSides = *f_sides
	c.IsPrint = *f_print
	c.IsStream = false

	c.Editor = *f_editor
	c.Number = *f_number
	c.MetaFile = *f_meta
	c.Files = files

	if !c.IsReview && !c.IsMemorize && !c.IsDone {
		c.IsReview = true
		c.IsMemorize = true
	}

	return c
}
