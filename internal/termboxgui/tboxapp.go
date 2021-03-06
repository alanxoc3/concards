package termboxgui

import (
	"github.com/alanxoc3/concards/internal"
	"github.com/alanxoc3/concards/internal/card"
	"github.com/alanxoc3/concards/internal/deck"
	"github.com/alanxoc3/concards/internal/file"
	"github.com/alanxoc3/concards/internal/history"
	"github.com/go-cmd/cmd"
	termbox "github.com/nsf/termbox-go"
)

func TermBoxRun(d *deck.Deck, cfg *internal.Config) error {
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
	statMode := false
	quitMode := false
	finishedEditing := false

	// Startup Hook
	cmd.NewCmd(cfg.EventStartupFile).Start()

	man.Save(d) // Save at beginning, and end of each editing command.
	for d.ReviewLen() > 0 {
		drawScreen(d, cardShown, helpMode, statMode, finishedEditing)
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
				statMode = false
			} else if inp == "s" {
				statMode = !statMode
				helpMode = false
			} else if inp == "w" {
				err = file.WritePredictsToFile(d.PredictList(), cfg.PredictFile)
				if err != nil {
					updateStatMsg(err.Error(), termbox.ColorRed)
				} else {
					err = file.WriteOutcomesToFile(d.OutcomeList(), cfg.OutcomeFile)
					if err != nil {
						updateStatMsg(err.Error(), termbox.ColorRed)
					} else {
						updateStatMsg("Cards were written.", termbox.ColorYellow)
					}
				}

			} else if !helpMode && !statMode {
				if inp == "1" {
					updateStatMsgAndCard(d, false, cfg)
					cardShown = 1
					man.Save(d)
				} else if inp == "2" {
					updateStatMsgAndCard(d, true, cfg)
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

func drawScreen(d *deck.Deck, cardShown int, helpMode, statMode, finishedEditing bool) {
	if finishedEditing {
		termbox.Sync()
	}

	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)

	if helpMode {
		displayHelpMode(termbox.ColorCyan)
	} else if statMode {
		displayStatMode(termbox.ColorYellow, d.TopPredict())
	} else {
		displayCardMode(d.TopCard(), cardShown)
	}

	tbprintStatusbar(d)
	tbprintStatMsg()
	flush()
}
