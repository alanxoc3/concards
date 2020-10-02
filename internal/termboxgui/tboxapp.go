package termboxgui

import (
	"github.com/alanxoc3/concards/internal/card"
	"github.com/alanxoc3/concards/internal/deck"
	"github.com/alanxoc3/concards/internal/file"
	"github.com/alanxoc3/concards/internal/history"
	termbox "github.com/nsf/termbox-go"
)

func TermBoxRun(d *deck.Deck, cfg *file.Config) error {
	err := termbox.Init()
	if err != nil {
		return err
	}

	defer termbox.Close()
	defer termbox.Sync() // TODO: Not sure what the purpose of this is.
	termbox.SetInputMode(termbox.InputEsc)
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
	flush()

	data = make([]byte, 0, 64)
	const coldef = termbox.ColorDefault

	man := history.NewManager()
	cardShown := 1
	helpMode := false
	quitMode := false
	finishedEditing := false

	man.Save(d) // Save at beginning, and end of each editing command.
	for d.ReviewLen() > 0 {
		drawScreen(d, helpMode, cardShown, finishedEditing)
		finishedEditing = false

		inp, err := updateInput()
		if err != nil {
			return err
		}

		if !quitMode {
			if inp == "q" {
				updateStatMsg("Quit concards? [y/N]", termbox.ColorYellow)
				quitMode = true
			} else if inp == "h" {
				helpMode = !helpMode
			} else if inp == "w" {
				err = file.WritePredictsToFile(d.PredictList(), cfg.PredictFile)
				if err != nil {
					updateStatMsg(err.Error(), termbox.ColorRed)
				} else {
					updateStatMsg("Cards were written.", termbox.ColorYellow)
				}
			} else if !helpMode {
				if inp == "1" {
					updateStatMsgAndCard(d, false)
					cardShown = 1
					man.Save(d)
				} else if inp == "2" {
					updateStatMsgAndCard(d, true)
					cardShown = 1
					man.Save(d)
				} else if inp == "d" {
					updateStatMsg("Deleted.", termbox.ColorYellow)
					cardShown = 1
					d.DropTop()
					man.Save(d)
				} else if inp == "e" {
					err := d.Edit(file.ReadCardsFromFile, func(filename string) ([]*card.Card, error) {
						return file.EditCards(filename, cfg)
					})

					if err != nil {
						updateStatMsg(err.Error(), termbox.ColorRed)
					} else {
						updateStatMsg("Card was successfully edited.", termbox.ColorYellow)
					}
					finishedEditing = true
					cardShown = 1
					man.Save(d)
				} else if inp == "u" {
					td, terr := man.Undo()
					if terr != nil {
						updateStatMsg(terr.Error(), termbox.ColorRed)
					} else {
						d.Clone(td)
						updateStatMsg("Undo.", termbox.ColorYellow)
						cardShown = 1
					}
				} else if inp == "r" {
					td, terr := man.Redo()
					if terr != nil {
						updateStatMsg(terr.Error(), termbox.ColorRed)
					} else {
						d.Clone(td)
						updateStatMsg("Redo.", termbox.ColorCyan)
						cardShown = 1
					}
				} else if inp == " " || inp == "\r" {
					cardShown++
					if l := d.TopCard().Len(); cardShown > l {
						cardShown = 1
					}
				}
			}
		} else {
			if inp == "y" {
				return nil
			} else {
				quitMode = false
				updateStatMsg("", coldef)
			}
		}

	}

	return nil
}

func drawScreen(d *deck.Deck, helpMode bool, cardShown int, finishedEditing bool) {
	if finishedEditing {
		termbox.Sync()
	}

	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)

	if helpMode {
		displayHelpMode(termbox.ColorCyan)
	} else {
		displayCardMode(d.TopCard(), cardShown)
	}

	tbprintStatusbar(d)
	tbprintStatMsg()
	flush()
}
