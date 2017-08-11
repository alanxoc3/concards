package termboxgui

import (
	"fmt"

	"github.com/alanxoc3/concards-go/algs"
	"github.com/alanxoc3/concards-go/card"
	"github.com/alanxoc3/concards-go/deck"
	runewidth "github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
)

var data []byte
var stat_msg string
var stat_msg_col termbox.Attribute

const coldef = termbox.ColorDefault

func update_input() (string, error) {
	current := ""

	// Resize the data.
	if cap(data)-len(data) < 32 {
		newdata := make([]byte, len(data), len(data)+32)
		copy(newdata, data)
		data = newdata
	}

	beg := len(data)
	d := data[beg : beg+32]
	if ev := termbox.PollRawEvent(d); ev.Type == termbox.EventRaw {
		data = data[:beg+ev.N]
		current = string(data)

		for {
			ev := termbox.ParseEvent(data)
			if ev.N == 0 {
				break
			}

			copy(data, data[ev.N:])
			data = data[:len(data)-ev.N]
		}
	} else if ev.Type == termbox.EventError {
		return "", ev.Err
	} // else, ignore it.

	return current, nil
}

func flush() {
	termbox.SetCursor(0, 0) // this line fixes a cursor glitch for after vim takes over the screen.
	termbox.HideCursor()
	termbox.Flush()
}

// ignores tabs, returns the final x and y position.
func tbprint(x, y int, fg, bg termbox.Attribute, msg string) (int, int) {
	xinitial := x
	for _, c := range msg {
		inc := runewidth.StringWidth(string(c))
		if char := string(c); char == "\n" {
			y++
			x = xinitial
		} else if char != "\t" {
			termbox.SetCell(x, y, c, fg, bg)
			x += inc
		}
	}

	return x, y
}

func tbprintwrap(x, y int, fg, bg termbox.Attribute, msg string) (int, int) {
	w, _ := termbox.Size()
	xinitial := x
	for _, c := range msg {
		inc := runewidth.StringWidth(string(c))
		if char := string(c); char == "\n" {
			y++
			x = xinitial
		} else if char != "\t" {
			termbox.SetCell(x, y, c, fg, bg)
			if x+inc > w {
				x = 0
				y++
			}

			x += inc
		}
	}

	return x, y
}

func tbhorizontal(y int, color termbox.Attribute) {
	w, _ := termbox.Size()

	for i := 0; i < w; i++ {
		termbox.SetCell(i, y, ' ', coldef, color)
	}
}

func tbvertical(x int, color termbox.Attribute) {
	_, h := termbox.Size()

	for i := 0; i < h; i++ {
		termbox.SetCell(x, i, ' ', coldef, color)
	}
}

// print question, returns final
func tbprint_gq(c *card.Card, color termbox.Attribute) (int, int) {
	grps := c.FormatGroups()
	_, y := tbprintwrap(0, 0, termbox.ColorCyan, coldef, grps)

	ques := fmt.Sprintf(c.FormatQuestion())
	return tbprint(0, y+2, color, coldef, ques)
}

// print question and answer
func tbprint_gqa(c *card.Card) {
	_, y := tbprint_gq(c, termbox.ColorCyan)

	ans := c.FormatAnswer()
	tbprintwrap(3, y+2, termbox.ColorWhite, coldef, ans)
}

func tbprint_statusbar(d deck.Deck) {
	_, h := termbox.Size()
	color := termbox.ColorBlue
	tbhorizontal(h-2, color)
	msg := fmt.Sprintf("concards - %d cards remain - type \"h\" for help", len(d))

	tbprint(0, h-2, termbox.ColorWhite|termbox.AttrBold, color, msg)
}

func display_help_mode(color termbox.Attribute) {
	str2 := "              Controls\n" +
		"------------------------------------\n" +
		"e, EDIT  - the current card\n" +
		"w, WRITE - your cards to their files\n" +
		"q, QUIT  - the program\n" +
		"h, HELP  - toggle this menu\n" +
		"\n" +
		"1, INPUT 1 - Not a clue.\n" +
		"2, INPUT 2 - Sounds familiar.\n" +
		"3, INPUT 3 - I know it!\n" +
		"\n" +
		"<space> or <enter> - reveal card\n"
	// 12 lines, longest line is 36 characters

	w, h := termbox.Size()

	// 36 characters wide, 12 lines tall.
	lw, lh := 36, 12

	x := w/2 - lw/2
	if x < 0 {
		x = 0
	}

	y := h/2 - lh/2
	if y < 0 {
		y = 0
	}

	tbprintwrap(x, y, color, coldef, str2)
}

func display_card_mode(c *card.Card, showAnswer bool) {
	if showAnswer {
		tbprint_gqa(c)
	} else {
		if c.HasAnswer() {
			tbprint_gq(c, termbox.ColorWhite)
		} else {
			tbprint_gq(c, termbox.ColorYellow)
		}
	}
}

func tbprint_stat_msg() {
	_, h := termbox.Size()
	color := termbox.ColorBlack
	tbhorizontal(h-1, color)

	tbprint(0, h-1, stat_msg_col, color, stat_msg)
}

func update_stat_msg_and_card(c *card.Card, k algs.Know) {
	if k == algs.NO {
		c.Metadata.Execute(algs.NO)
		update_stat_msg("Not a clue, card put to the back of the pile.", termbox.ColorRed)
	} else if k == algs.IDK {
		c.Metadata.Execute(algs.IDK)
		update_stat_msg("Sounds familiar, card put to the back of the pile.", termbox.ColorYellow)
	} else if k == algs.YES {
		c.Metadata.Execute(algs.YES)
		time := c.Metadata.Next.Format("Mon Jan 2, 2006 @ 15:04")
		update_stat_msg(fmt.Sprintf("I know it! Next review is %s.", time), termbox.ColorCyan)
	}
}

func update_stat_msg(msg string, color termbox.Attribute) {
	stat_msg = msg
	stat_msg_col = color
}
