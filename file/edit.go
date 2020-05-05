package file

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/alanxoc3/concards/core"
)

// Assumes the deck is sorted how you want it to be sorted.
func EditFile(d *core.Deck, cfg *Config) error {
   h, c, _ := d.Top()
   if c == nil {
		return fmt.Errorf("Error: The deck is empty.")
   }

   f := c.File
	cmd := exec.Command(cfg.Editor, f)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Error: The editor returned an error code.")
	}

   nd := core.NewDeck()
   for k, v := range d.Mmap { nd.Mmap[k] = v }
   ReadCardsToDeck(nd, f)

   if !cfg.IsMemorize { nd.FilterOutMemorize() }
   if !cfg.IsReview { nd.FilterOutReview() }
   if !cfg.IsDone { nd.FilterOutDone() }

   d.FilterOutFileDeck(f, nd)
   card_index := 0
   if h == d.TopHash() {
      card_index = 1
   }

   for i := nd.Len() - 1; i >= 0; i-- {
      d.InsertCard(nd.GetCard(i), card_index)
   }

	return nil
}
