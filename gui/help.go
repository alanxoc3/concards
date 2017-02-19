package gui

import "fmt"

func (cfg * Config) Usage() (ret string) {
	ret += fmt.Sprintf("Usage:\n")
	ret += fmt.Sprintf("  concards [OPTIONS...] FILE1\n")
	ret += fmt.Sprintf("  concards [OPTIONS...] FILE1 FILE2...\n")
	return
}

func (cfg * Config) Help() (ret string) {
	opts := cfg.Opts
	ret = cfg.Usage()

	ret += "\n" + fmt.Sprintf("Card Limiting Options (may use multiple):\n")
	for i := 0; i <= 6; i++ { ret += optionToString(opts[i]) }

	ret += "\n" + fmt.Sprintf("Performance Mode Options (may use one):\n")
	for i := 7; i <= 9; i++ { ret += optionToString(opts[i]) }

	ret += "\n" + fmt.Sprintf("Other Options:\n")
	for i := 10; i <= 15; i++ { ret += optionToString(opts[i]) }

	return ret
}

func GenOptions() []*Option {
	var opts []*Option

	opts = append(opts, newOptionNoParam("r", "review", "Show cards available to be reviewed."))
	opts = append(opts, newOptionNoParam("m", "memorize", "Show cards available to be memorized."))
	opts = append(opts, newOptionNoParam("d", "done", "Show cards not available to be reviewed or memorized."))
	opts = append(opts, newOption("g", "groups", "grp", "Limit the cards in the program to only within the groups of 'grp'."))
	opts = append(opts, newOption("n", "number", "#", "Limit the number of cards in the program to 'p'."))
	opts = append(opts, newOptionNoParam("o", "one", "Limit the number of cards to only one card. Same as '-n 1'."))
	opts = append(opts, newOptionNoParam("", "no-main", "Disables the main screen. You may use in combination with any of the above."))

	opts = append(opts, newOptionNoParam("e", "edit", "Edit the cards with the default editor instead of reviewing them."))
	opts = append(opts, newOptionNoParam("p", "print", "Print out what the output file would be, based on the input files."))
	opts = append(opts, newOptionNoParam("u", "update", "Simply updates the file to the correct formatting."))

	opts = append(opts, newOptionNoParam("h", "help", "Prints out a usage/help menu."))
	opts = append(opts, newOptionNoParam("v", "version", "Prints out which version is being used."))
	opts = append(opts, newOptionNoParam("", "color", "Enable the cards to have color. Disabled by default."))
	opts = append(opts, newOptionNoParam("", "no-write", "Concards will not write to any file."))
	opts = append(opts, newOption("", "editor", "e", "Change the editor 'ed' used when editing. Default is \"$EDITOR\"."))
	opts = append(opts, newOption("", "def-alg", "a", "The default algorithm 'alg' used. This is set to 'SM2' by default."))

	return opts
}

/*
// Expected Output
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
