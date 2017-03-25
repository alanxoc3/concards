package gui

const (
	REVIEW   = iota
	MEMORIZE = iota
	DONE     = iota
	GROUPS   = iota
	NUMBER   = iota
	ONE      = iota
	EDIT     = iota
	PRINT    = iota
	UPDATE   = iota
	HELP     = iota
	VERSION  = iota
	COLOR    = iota
	NOMAIN   = iota
	NOWRITE  = iota
	EDITOR   = iota
	DEFALG   = iota
)

func (cfg *Config) Usage() (ret string) {
	ret += "Usage:\n"
	ret += "   concards [OPTIONS...] FILE1\n"
	ret += "   concards [OPTIONS...] DIRECTORY\n"
	ret += "   concards [OPTIONS...] FILE1 FILE2...\n"
	return
}

func (cfg *Config) Help() (ret string) {
	ret = cfg.Usage()
	ret += "\n"
	ret += "For more detailed options, read the fine man page.\n"

	return ret
}

func genOptions() []*Option {
	var opts []*Option

	opts = append(opts, newOptionNoParam('r', "review", "Show cards available to be reviewed."))
	opts = append(opts, newOptionNoParam('m', "memorize", "Show cards available to be memorized."))
	opts = append(opts, newOptionNoParam('d', "done", "Show cards not available to be reviewed or memorized."))
	opts = append(opts, newOption('g', "groups", "grp", "Limit the cards in the program to only within the groups of 'grp'."))
	opts = append(opts, newOption('n', "number", "#", "Limit the number of cards in the program to 'p'."))
	opts = append(opts, newOptionNoParam('o', "one", "Limit the number of cards to only one card. Same as '-n 1'."))

	opts = append(opts, newOptionNoParam('e', "edit", "Edit the cards with the default editor instead of reviewing them."))
	opts = append(opts, newOptionNoParam('p', "print", "Print out what the output file would be, based on the input files."))
	opts = append(opts, newOptionNoParam('u', "update", "Simply updates the file to the correct formatting."))

	opts = append(opts, newOptionNoParam('h', "help", "Prints out a usage/help menu."))
	opts = append(opts, newOptionNoParam('v', "version", "Prints out which version is being used."))
	opts = append(opts, newOptionNoParam(0, "color", "Enable the cards to have color. Disabled by default."))
	opts = append(opts, newOptionNoParam(0, "no-main", "Disables the main screen."))
	opts = append(opts, newOptionNoParam(0, "no-write", "Concards will not write to any file."))
	opts = append(opts, newOption(0, "editor", "e", "Change the editor 'ed' used when editing. Default is \"$EDITOR\"."))
	opts = append(opts, newOption(0, "def-alg", "a", "The default algorithm 'alg' used. This is set to 'SM2' by default."))

	return opts
}
