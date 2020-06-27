package termboxgui

import (
	"github.com/alanxoc3/concards/core"
	"github.com/alanxoc3/concards/file"
	termbox "github.com/nsf/termbox-go"
)

func TermBoxRun(d *core.Deck, cfg *file.Config) error {
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

	cardShown := 1
	helpMode := false
	quitMode := false
	finishedEditing := false

	save(d) // Save at beginning, and end of each editing command.
	for d.Len() > 0 {
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
				err = file.WriteMetasToFile(d, cfg.MetaFile)
				if err != nil {
					updateStatMsg(err.Error(), termbox.ColorRed)
				} else {
					updateStatMsg("Cards were written.", termbox.ColorYellow)
				}
			} else if !helpMode {
				if inp == "1" {
					updateStatMsgAndCard(d, core.NO)
					cardShown = 1
					d.TopTo(3)
					save(d)
				} else if inp == "2" {
					updateStatMsgAndCard(d, core.IDK)
					cardShown = 1
					d.TopTo(6)
					save(d)
				} else if inp == "3" {
					updateStatMsgAndCard(d, core.YES)
					cardShown = 1
					d.DelTop()
					save(d)
				} else if inp == "d" {
					updateStatMsg("Deleted.", termbox.ColorYellow)
					cardShown = 1
					d.DelTop()
					save(d)
				} else if inp == "f" {
					updateStatMsg("Forgotten.", termbox.ColorYellow)
					cardShown = 1
					d.ForgetTop()
					save(d)
				} else if inp == "s" {
					updateStatMsg("Skipped.", termbox.ColorYellow)
					cardShown = 1
					d.TopToEnd()
					save(d)
				} else if inp == "e" {
					err := file.EditFile(d, cfg, file.ReadCards, file.EditCards)

					if err != nil {
						updateStatMsg(err.Error(), termbox.ColorRed)
					} else {
						updateStatMsg("Card was successfully edited.", termbox.ColorYellow)
					}
					finishedEditing = true
					cardShown = 1
					save(d)
				} else if inp == "u" {
					td, terr := undo()
					if terr != nil {
						updateStatMsg(terr.Error(), termbox.ColorRed)
					} else {
						d.Clone(td)
						updateStatMsg("Undo.", termbox.ColorYellow)
						cardShown = 1
					}
				} else if inp == "r" {
					td, terr := redo()
					if terr != nil {
						updateStatMsg(terr.Error(), termbox.ColorRed)
					} else {
						d.Clone(td)
						updateStatMsg("Redo.", termbox.ColorCyan)
						cardShown = 1
					}
				} else if inp == " " || inp == "\r" {
					cardShown++
					if l := d.GetCard(0).Len(); cardShown > l {
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

func drawScreen(d *core.Deck, helpMode bool, cardShown int, finishedEditing bool) {
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
