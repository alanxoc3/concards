package termhelp

func Version() (vers string) {
	return "concards v1.1"
}

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


/*
  -p  --print     Prints all cards, slightly formatted.
  -h  --help      If you need assistance.
  -v  --version   Which version are you on again?
  */

func genOptions() []*Option {
	var opts []*Option

	opts = append(opts, newOptionNoParam('r', "review", "Show cards available to be reviewed."))
	opts = append(opts, newOptionNoParam('m', "memorize", "Show cards available to be memorized."))
	opts = append(opts, newOptionNoParam('d', "done", "Show cards not available to be reviewed or memorized."))
	opts = append(opts, newOption('n', "number", "#", "Only process \"#\" cards."))
	opts = append(opts, newOptionNoParam('p', "print", "Prints all cards, one line per card."))
	opts = append(opts, newOptionNoParam('h', "help", "If you need assistance."))
	opts = append(opts, newOptionNoParam('v', "version", "Which concards version is this?"))
	opts = append(opts, newOption('E', "editor", "e", "Which editor concards should use. Defaults to \"$EDITOR\"."))
	opts = append(opts, newOption('M', "meta", "f", "Location of concards meta file. Defaults to "$CONCARDS_META" or \"~/.concards-meta\"."))

	return opts
}
