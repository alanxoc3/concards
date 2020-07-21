package termboxgui

import (
	"fmt"

	"github.com/alanxoc3/concards/core"
	runewidth "github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
)

var data []byte
var statMsg string
var statMsgCol termbox.Attribute

const coldef = termbox.ColorDefault

func updateInput() (string, error) {
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

func tbPrintHelper(x, y int, fg, bg termbox.Attribute, msg string, wrap bool) (int, int) {
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
	return tbPrintHelper(x, y, fg, bg, msg, false)
}

// returns final x and y position.
func tbprintwrap(x, y int, fg, bg termbox.Attribute, msg string) (int, int) {
	return tbPrintHelper(x, y, fg, bg, msg, true)
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

func tbprintCard(c *core.Card, amount int) {
	y := 0

	for i := 0; i < c.Len() && i < amount; i++ {
		color := termbox.ColorCyan
		if i > 0 {
			color = termbox.ColorWhite
		}
		_, y = tbprintwrap(0, y, color, coldef, c.GetFactEsc(i))
		y++
	}
}

func tbprintStatusbar(d *core.Deck) {
	_, h := termbox.Size()
	color := termbox.ColorBlue
	tbhorizontal(h-2, color)
	msg := fmt.Sprintf("%d cards - %s", d.Len(), d.TopCard().GetFile())

	tbprint(0, h-2, termbox.ColorWhite|termbox.AttrBold, color, msg)
}

func displayHelpMode(color termbox.Attribute) {
	str2 := "              Controls\n" +
		"------------------------------------\n" +
		"[d]elete: Remove card from session.\n" +
		"[e]dit:   Open card in editor.\n" +
		"[f]orget: Reset card's progress.\n" +
		"[h]elp:   Toggle this menu.\n" +
		"[k]now:   Exit the program.\n" +
		"[q]uit:   Exit the program.\n" +
		"[r]edo:   Redo the undo.\n" +
		"[s]kip:   Put the card at the end.\n" +
		"[u]ndo:   Undo last action.\n" +
		"[w]rite:  Write state to meta file.\n" +
		"\n" +
		"[1]: Not a clue.\n" +
      "[2]: Sounds familiar.\n" +
      "[3]: I know it!\n" +
		"\n" +
      "[space,enter]: Reveal next side.\n"
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

func displayCardMode(c *core.Card, showAnswer int) {
	tbprintCard(c, showAnswer)
}

func tbprintStatMsg() {
	_, h := termbox.Size()
	color := termbox.ColorBlack
	tbhorizontal(h-1, color)

	tbprint(0, h-1, statMsgCol, color, statMsg)
}

func updateStatMsgAndCard(d *core.Deck, k core.Know) {
	h, _, m := d.Top()
	if m == nil {
		m = core.NewDefaultMeta("sm2")
	}

	if k == core.NO {
		m, _ = m.Exec(core.NO)
		updateStatMsg("No! Try again soon.", termbox.ColorRed)
	} else if k == core.IDK {
		m, _ = m.Exec(core.IDK)
		updateStatMsg("Idk! Try again in a bit.", termbox.ColorYellow)
	} else if k == core.YES {
		m, _ = m.Exec(core.YES)
		time := m.Next.Format("Mon 2 Jan 2006 @ 15:04")
		updateStatMsg(fmt.Sprintf("Yes! Next review is %s.", time), termbox.ColorCyan)
	} else if k == core.KNOW {
		m, _ = m.Exec(core.KNOW)
		updateStatMsg(fmt.Sprintf("Known! Next review is %d years from now.", core.YearsToAddForKnown), termbox.ColorCyan)
	}

	d.AddMeta(h, m)
}

func updateStatMsg(msg string, color termbox.Attribute) {
	statMsg = msg
	statMsgCol = color
}
