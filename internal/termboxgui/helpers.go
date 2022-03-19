package termboxgui

import (
	"fmt"
	"time"
	"github.com/go-cmd/cmd"

	"github.com/alanxoc3/concards/internal"
	"github.com/alanxoc3/concards/internal/card"
	"github.com/alanxoc3/concards/internal/deck"
	"github.com/alanxoc3/concards/internal/meta"
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

func tbprintCard(c *card.Card, amount int) {
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

func tbprintStatusbar(d *deck.Deck) {
	_, h := termbox.Size()
	color := termbox.ColorBlue
	tbhorizontal(h-1, color)
	msg := fmt.Sprintf("%s", d.TopCard().File())

	tbprint(0, h-1, termbox.ColorWhite|termbox.AttrBold, color, msg)
}

func displayHelpMode(color termbox.Attribute) {
	helpStr := `[d]elete: Remove card from session.
[e]dit:   Open card in editor.
[h]elp:   Toggle this menu.
[q]uit:   Exit the program.
[r]edo:   Redo the undo.
[s]tat:   Show card statistics.
[u]ndo:   Undo last action.
[w]rite:  Write state to meta file.

[1]: No!
[2]: Yes!

[space,enter]: Reveal next side.`
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

	tbprint(x, y, color, coldef, helpStr)
}

func displayStatMode(color termbox.Attribute, p *meta.Predict) {
	statStr := fmt.Sprintf(`
Hash: %s
Alg:  %s

Total Yes Count: %d
Total No Count:  %d
Current Streak:  %d

Last Reviewed:  %s
Planned Review: %s

Press [s] to leave this screen.
`, p.Hash().String(), p.Name(), p.YesCount(), p.NoCount(), p.Streak(), p.Curr().Local().Format(time.RFC1123), p.Next().Local().Format(time.RFC1123))

	w, h := termbox.Size()
	h = h - 2 // Status bar at the bottom.

	// characters wide, lines tall.
	lw, lh := 37, 17

	x := w/2 - lw/2

	y := h/2 - lh/2
	if y < 0 {
		y = 0
	}

	tbprint(x, y, color, coldef, statStr)
}

func displayCardMode(c *card.Card, showAnswer int) {
	tbprintCard(c, showAnswer)
}

func tbprintStatMsg() {
	_, h := termbox.Size()
	color := termbox.ColorBlack
	tbhorizontal(h-2, color)

	tbprint(0, h-2, statMsgCol, color, statMsg)
}

func updateStatMsgAndCard(d *deck.Deck, input bool, cfg *internal.Config) {
	m, err := d.ExecTop(input, time.Now())

	if err != nil {
		updateStatMsg("Problem reading the card :(.", termbox.ColorRed)
	} else {
    		// Review Hook
		cmd.NewCmd(cfg.EventReviewFile).Start()

		time := m.Next().Local().Format("Mon 2 Jan 2006 @ 15:04")

		if input {
			updateStatMsg(fmt.Sprintf("Yes! Next review is %s.", time), termbox.ColorCyan)
		} else {
			updateStatMsg(fmt.Sprintf("No! Next review is %s.", time), termbox.ColorRed)
		}
	}
}

func updateStatMsg(msg string, color termbox.Attribute) {
	statMsg = msg
	statMsgCol = color
}
