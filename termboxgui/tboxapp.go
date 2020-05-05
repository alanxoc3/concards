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

	card_shown := 1
	help_mode := false
	quit_mode := false
	finished_editing := false

	save(d) // Save at beginning, and end of each editing command.
	for d.Len() > 0 {
		draw_screen(d, help_mode, card_shown, finished_editing)
		finished_editing = false

		inp, err := update_input()
		if err != nil {
			return err
		}

		if !quit_mode {
			if inp == "q" {
				update_stat_msg("Quit concards? [y/N]", termbox.ColorYellow)
				quit_mode = true
			} else if inp == "h" {
				help_mode = !help_mode
			} else if inp == "w" {
            err = file.WriteMetasToFile(d, cfg.MetaFile)
				if err != nil {
					update_stat_msg(err.Error(), termbox.ColorRed)
				} else {
					update_stat_msg("Cards were written.", termbox.ColorYellow)
				}
			} else if !help_mode {
				if inp == "1" {
					update_stat_msg_and_card(d, core.NO)
					card_shown = 1
               d.TopTo(3)
					save(d)
				} else if inp == "2" {
					update_stat_msg_and_card(d, core.IDK)
					card_shown = 1
               d.TopTo(6)
					save(d)
				} else if inp == "3" {
					update_stat_msg_and_card(d, core.YES)
					card_shown = 1
					d.DelTop()
					save(d)
				} else if inp == "d" {
               update_stat_msg("Deleted.", termbox.ColorYellow)
					card_shown = 1
					d.DelTop()
					save(d)
				} else if inp == "f" {
               update_stat_msg("Forgotten.", termbox.ColorYellow)
					card_shown = 1
					d.ForgetTop()
					save(d)
				} else if inp == "s" {
               update_stat_msg("Skipped.", termbox.ColorYellow)
					card_shown = 1
               d.TopToEnd()
					save(d)
				} else if inp == "e" {
					err := file.EditFile(d, cfg)

					if err != nil {
						update_stat_msg(err.Error(), termbox.ColorRed)
					} else {
						update_stat_msg("Card was successfully edited.", termbox.ColorYellow)
					}
					finished_editing = true
					card_shown = 1
					save(d)
				} else if inp == "u" {
					td, terr := undo()
					if terr != nil {
						update_stat_msg(terr.Error(), termbox.ColorRed)
					} else {
                  d.Clone(td)
						update_stat_msg("Undo.", termbox.ColorYellow)
						card_shown = 1
					}
				} else if inp == "r" {
					td, terr := redo()
					if terr != nil {
						update_stat_msg(terr.Error(), termbox.ColorRed)
					} else {
                  d.Clone(td)
						update_stat_msg("Redo.", termbox.ColorCyan)
						card_shown = 1
					}
				} else if inp == " " || inp == "\r" {
               card_shown++
               if l := len(d.GetCard(0).Facts); card_shown > l {
                  card_shown = 1
               }
				}
			}
		} else {
			if inp == "y" {
				return nil
			} else {
				quit_mode = false
				update_stat_msg("", coldef)
			}
		}

	}

	return nil
}

func draw_screen(d *core.Deck, help_mode bool, card_shown int, finished_editing bool) {
	if finished_editing {
		termbox.Sync()
	}

	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)

	if help_mode {
		display_help_mode(termbox.ColorCyan)
	} else {
		display_card_mode(d.TopCard(), card_shown)
	}

	tbprint_statusbar(d)
	tbprint_stat_msg()
	flush()
}
