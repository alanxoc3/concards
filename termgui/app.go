package termgui

import (
	"fmt"
	"io"
	"strings"

	"github.com/alanxoc3/concards-go/termhelp"
	"github.com/chzyer/readline"
)

var completer = readline.NewPrefixCompleter(
	readline.PcItem("reverse"),
	readline.PcItem("reshow"),
	readline.PcItem("help"),
)

func Run() {
	// Some basic Readline setup
	l, err := readline.NewEx(&readline.Config{
		Prompt:          "> ",
		HistoryFile:     "/tmp/concards.tmp",
		AutoComplete:    completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",

		HistorySearchFold: true,
	})

	if err != nil {
		panic(err)
	}

	defer l.Close()

	setPasswordCfg := l.GenPasswordConfig()
	setPasswordCfg.SetListener(func(line []rune, pos int, key rune) (newLine []rune, newPos int, ok bool) {
		l.SetPrompt(fmt.Sprintf("Enter password(%v): ", len(line)))
		l.Refresh()
		return nil, 0, false
	})

	for {
		// Read line, quit on a ctrl-d
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		// Some basic command line parameters.
		line = strings.TrimSpace(line)
		switch {
		case line == "help":
			fmt.Println(termhelp.Help())
		case strings.HasPrefix(line, "settop"):
			readline.TopText = line + "\n"
		case strings.HasPrefix(line, "setprompt"):
			if len(line) <= 10 {
				break
			}
		}
	}
}
