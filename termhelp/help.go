package termhelp

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
	EDITOR   = iota
)

func Usage() (ret string) {
	ret += "Usage: concards [OPTION]... [FILE]...\n"
	return
}

func Help() (ret string) {
	ret = Usage()
	ret += "\nOptions:\n"

	opts := genOptions()
	for _, opt := range opts {
		ret += opt.ToString()
	}

	ret += "\nFor more detailed options, read the fine man page."

	return ret
}

func genOptions() []*Option {
	var opts []*Option

	opts = append(opts, newOptionNoParam('r', "review", "Show cards available to be reviewed."))
	opts = append(opts, newOptionNoParam('m', "memorize", "Show cards available to be memorized."))
	opts = append(opts, newOptionNoParam('d', "done", "Show cards not available to be reviewed or memorized."))
	opts = append(opts, newOption('g', "groups", "grps", "Limits the cards to only those a part of one of the groups in \"grps\"."))
	opts = append(opts, newOption('n', "number", "#", "Limit the number of cards in the program to \"#\"."))
	opts = append(opts, newOptionNoParam('o', "one", "Limit the number of cards to only one card. Same as \"-n 1\"."))

	opts = append(opts, newOptionNoParam('e', "edit", "Edit the cards with the default editor instead of reviewing them."))
	opts = append(opts, newOptionNoParam('p', "print", "Print what the combined file of all the cards would be."))
	opts = append(opts, newOptionNoParam('u', "update", "Simply updates the files to the correct formatting."))

	opts = append(opts, newOptionNoParam('h', "help", "Prints out a usage/help menu."))
	opts = append(opts, newOptionNoParam('v', "version", "Prints out which version is being used."))
	opts = append(opts, newOption(0, "editor", "e", "Change the editor \"e\" used when editing. Default is \"$EDITOR\"."))

	return opts
}
