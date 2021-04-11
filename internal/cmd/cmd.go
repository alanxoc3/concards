package cmd

import (
	"os"
	"os/user"
)

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

/*
func GenConfig(version string) *Config {
	// Create new parser object
	parser := argparse.NewParser("concards", "Concards is a simple CLI based SRS flashcard app.")

	// Create flags
	fVersion := parser.Flag("V", "version", &argparse.Options{Help: "Concards build information"})
	fReview := parser.Flag("r", "review", &argparse.Options{Help: "Show cards available to be reviewed"})
	fMemorize := parser.Flag("m", "memorize", &argparse.Options{Help: "Show cards available to be memorized"})
	fDone := parser.Flag("d", "done", &argparse.Options{Help: "Show cards not available to be reviewed or memorized"})
	fPrint := parser.Flag("p", "print", &argparse.Options{Help: "Print all cards, one card per line"})
	fFileList := parser.Flag("l", "files-with-cards", &argparse.Options{Help: "Print the file paths that have cards"})
	fNumber := parser.Int("n", "number", &argparse.Options{Default: 0, Help: "How many cards to review"})
	fEditor := parser.String("E", "editor", &argparse.Options{Default: defaultEditor(), Help: "Defaults to \"$EDITOR\" or \"vi\""})
	fPredictFile := parser.String("P", "predict", &argparse.Options{
      Default: defaultEnv("CONCARDS_PREDICT", "predict"),
      Help: "Defaults to \"$CONCARDS_PREDICT\" or \"~/.config/concards/predict\"" })
	fOutcomeFile := parser.String("O", "outcome", &argparse.Options{
      Default: defaultEnv("CONCARDS_OUTCOME", "outcome"),
      Help: "Defaults to \"$CONCARDS_OUTCOME\" or \"~/.config/concards/outcome\"" })

	parser.HelpFunc = func(c *argparse.Command, msg interface{}) string {
		var help string
		help += fmt.Sprintf("%s\n\nUsage:\n  %s [OPTIONS]... [FILE|FOLDER]...\n\nOptions:\n", c.GetDescription(), c.GetName())

		for _, arg := range c.GetArgs() {
			if arg.IsFlag() {
				help += fmt.Sprintf("  -%s  --%-17s %s.\n", arg.GetSname(), arg.GetLname(), arg.GetOpts().Help)
			} else {
				help += fmt.Sprintf("  -%s  --%-17s %s.\n", arg.GetSname(), arg.GetLname()+" "+arg.GetSname(), arg.GetOpts().Help)
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

	if *fVersion {
		fmt.Printf("concards %s\n", version)
		os.Exit(0)
	}

   if *fPredictFile == "" { internal.AssertError("No predict file available.") }
   if *fOutcomeFile == "" { internal.AssertError("No outcome file available.") }

	c := &Config{}

	c.IsReview = *fReview
	c.IsMemorize = *fMemorize
	c.IsDone = *fDone
	c.IsPrint = *fPrint
	c.IsFileList = *fFileList
	c.IsStream = false

	c.Editor = *fEditor
	c.Number = *fNumber
	c.PredictFile = *fPredictFile
	c.OutcomeFile = *fOutcomeFile
	c.Files = files

	if !c.IsReview && !c.IsMemorize && !c.IsDone {
		c.IsReview = true
		c.IsMemorize = true
		c.IsDone = true
	}

	return c
}
*/
