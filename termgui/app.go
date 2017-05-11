package termgui

import (
	"fmt"
	"io"
	"strings"

	"github.com/alanxoc3/concards-go/algs"
	"github.com/alanxoc3/concards-go/deck"
	"github.com/alanxoc3/concards-go/termhelp"
	"github.com/chzyer/readline"
)

var completer = readline.NewPrefixCompleter(
	readline.PcItem("yes"),
	readline.PcItem("no"),
	readline.PcItem("idk"),
	readline.PcItem("editdeck"),
	readline.PcItem("editcard"),
	readline.PcItem("ask"),
	readline.PcItem("help"),
	readline.PcItem("show"),
	readline.PcItem("grps"),

	// temp command
	readline.PcItem("detail"),
	readline.PcItem("howmany"),
)

func Run(d deck.Deck, cfg *termhelp.Config) {
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
			fmt.Println("help - displays this menu.              ")
			fmt.Println("ask  - shows the question.              ")
			fmt.Println("show - shows the answer.                ")
			fmt.Println("grps - shows the groups this card is in.")
			fmt.Println("yes  - passes the card.                 ")
			fmt.Println("no   - fails the card.                  ")
			fmt.Println("idk  - fails the card, but less harshly.")
		case line == "ask":
			fmt.Println(d.Top().FormatQuestion())
		case line == "show":
			fmt.Println(d.Top().FormatAnswer())
		case line == "grps":
			fmt.Println(d.Top().FormatGroups())
		case line == "yes":
			d.Top().Metadata.Execute(algs.YES)
			d = d[1:] // pop top
		case line == "no":
			d.Top().Metadata.Execute(algs.NO)
			d = append(d[1:], d[0]) // top to bottom
		case line == "idk":
			d.Top().Metadata.Execute(algs.IDK)
			d = append(d[1:], d[0]) // top to bottom

		case line == "editdeck":
			deck.EditDeck(cfg.Editor, d, "You may ONLY EDIT the cards here.\nREARRANGING, DELETING, or ADDING cards WILL CORRUPT your files.")

		case line == "editcard":
			deck.EditCard(cfg.Editor, d.Top(), "You may ONLY EDIT the cards here.\nREARRANGING, DELETING, or ADDING cards WILL CORRUPT your files.")

		case line == "detail":
			fmt.Println(d.Top().FormatFile())
		case line == "howmany":
			fmt.Printf("There are %d cards left.\n", d.Len())
		}

		if d.Top() == nil {
			return
		}
	}
}
