package gui

type Config struct {
	// The various true or false options
	numberEnabled bool
	groupsEnabled bool

	noOutputFile bool
	version      bool
	help         bool
	color        bool
	mainScreen   bool

	review   bool
	memorize bool
	done     bool

	// The variable options passed in.
	number int
	groups map[string]bool

	Opts []*Option
}

func GenConfig() *Config {
	var cfg Config
	cfg.Opts = GenOptions()

	return &cfg
}

/*
Usage:
  concards [OPTIONS...] FILE1
  concards [OPTIONS...] FILE1 FILE2...

Card Limiting Options (may use multiple):
  -r, --review          Show cards available to be reviewed.
  -m, --memorize        Show cards available to be memorized.
  -d, --done            Show cards not available to be reviewed or memorized.
  -g, --groups grp      Limit the cards in the program to only the groups of p.
  -n, --number num      Limit the number of cards in the program to 'p'.
  -o, --one             Limit the number of cards to only one card. Same as '-n 1'.
      --no-main         Disables the main screen. You may use in combination with any of the above.

Performance Mode Options (may use one):
  -e, --edit            Edit the cards with the default editor instead of reviewing them.
  -p, --print           Print out what the output file would be, based on the input files.
  -u, --update          Simply updates the file to the correct formatting.

Other Options:
  -h, --help            Prints out a usage/help menu.
  -v, --version         Prints out which version is being used.
      --color           Enable the cards to have color. Disabled by default.
      --no-write        Concards will not write to any file.
      --editor ed       Change the editor 'ed' used when editing. Default is "$EDITOR".
      --def-alg alg     The default algorithm 'alg' used. This is set to 'SM2' by default.

*/
