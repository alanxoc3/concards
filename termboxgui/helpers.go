package termboxgui

import (
	"fmt"

	"github.com/alanxoc3/concards/core"
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

func _tb_print_helper(x, y int, fg, bg termbox.Attribute, msg string, wrap bool) (int, int) {
	w, _ := termbox.Size()
	xinitial := x
	for _, c := range msg {
		inc := runewidth.StringWidth(string(c))
		char := string(c)

		if (wrap && x+inc > w) || char == "\n" {
			x = xinitial
			y++
		}

		termbox.SetCell(x, y, c, fg, bg)
		x += inc
	}

	return x, y
}

// ignores tabs, returns the final x and y position.
func tbprint(x, y int, fg, bg termbox.Attribute, msg string) (int, int) {
	return _tb_print_helper(x,y,fg,bg,msg,false)
}

// returns final x and y position.
func tbprintwrap(x, y int, fg, bg termbox.Attribute, msg string) (int, int) {
	return _tb_print_helper(x,y,fg,bg,msg,true)
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

func tbprint_card(c *core.Card, amount int) {
   y := 0

   for i := 0; i < len(c.Facts) && i < amount; i++ {
      color := termbox.ColorCyan
      if i > 0 { color = termbox.ColorWhite }
      _, y = tbprintwrap(0, y, color, coldef, c.Facts[i])
      y++
   }
}

func tbprint_statusbar(d *core.Deck) {
	_, h := termbox.Size()
	color := termbox.ColorBlue
	tbhorizontal(h-2, color)
	msg := fmt.Sprintf("%d concards - %s", d.Len(), d.TopCard().File)

	tbprint(0, h-2, termbox.ColorWhite|termbox.AttrBold, color, msg)
}

func display_help_mode(color termbox.Attribute) {
	str2 := "              Controls\n" +
		"------------------------------------\n" +
		"e, EDIT   - the current card\n" +
		"d, DELETE - delete the current card\n" +
		"s, SKIP   - skips the current card\n" +
		"f, FORGET - removes card's progress\n" +
		"w, WRITE  - your cards to their files\n" +
		"q, QUIT   - the program\n" +
		"h, HELP   - toggle this menu\n" +
		"u, UNDO   - you messed up\n" +
		"r, REDO   - maybe not\n" +
		"\n" +
		"1, INPUT 1 - Not a clue.\n" +
		"2, INPUT 2 - Sounds familiar.\n" +
		"3, INPUT 3 - I know it!\n" +
		"\n" +
		"<space> or <enter> - reveal card\n"
	// 12 lines, longest line is 36 characters

	w, h := termbox.Size()
	h = h - 2 // Status bar at the bottom.

	// characters wide, lines tall.
	lw, lh := 37, 17

	x := w/2 - lw/2

	y := h/2 - lh/2
	if y < 0 {
		y = 0
	}

	tbprint(x, y, color, coldef, str2)
}

func display_card_mode(c *core.Card, showAnswer int) {
   tbprint_card(c, showAnswer)
}

func tbprint_stat_msg() {
	_, h := termbox.Size()
	color := termbox.ColorBlack
	tbhorizontal(h-1, color)

	tbprint(0, h-1, stat_msg_col, color, stat_msg)
}

func update_stat_msg_and_card(d *core.Deck, k core.Know) {
   h, _, m := d.Top()
   if m == nil {
      m = core.NewDefaultMeta("sm2")
   }

	if k == core.NO {
      m, _ = m.Exec(core.NO)
      update_stat_msg("Not a clue, card put to the back of the pile.", termbox.ColorRed)
	} else if k == core.IDK {
      m, _ = m.Exec(core.IDK)
      update_stat_msg("Sounds familiar, card put to the back of the pile.", termbox.ColorYellow)
	} else if k == core.YES {
      m, _ = m.Exec(core.YES)
      time := m.Next.Format("Mon 2 Jan 2006 @ 15:04")
      update_stat_msg(fmt.Sprintf("I know it! Next review is %s.", time), termbox.ColorCyan)
	}

   d.AddMeta(h, m)
}

func update_stat_msg(msg string, color termbox.Attribute) {
	stat_msg = msg
	stat_msg_col = color
}
