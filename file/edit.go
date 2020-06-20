package file

import (
   "fmt"
   "os"
   "os/exec"

   "github.com/alanxoc3/concards/core"
)

// Assumes the deck is sorted how you want it to be sorted.
func EditFile(d *core.Deck, cfg *Config) error {
   if d.IsEmpty() {
      return fmt.Errorf("Error: The deck is empty.")
   }

   // We need to get information for the top card first.
   cur_hash, cur_card, cur_meta := d.Top()
   file_name := cur_card.GetFile()

   // Save the contents of the file now.
   deck_before := core.NewDeck()
   ReadCardsToDeck(deck_before, file_name)

   // Then edit the file.
   cmd := exec.Command(cfg.Editor, file_name)
   cmd.Stdin = os.Stdin
   cmd.Stdout = os.Stdout

   if err := cmd.Run(); err != nil {
      return fmt.Errorf("Error: The editor returned an error code.")
   }

   // Save the contents of the file after.
   deck_after := core.NewDeck()
   ReadCardsToDeck(deck_after, file_name)

   // Take out any card that was removed from the file.
   d.FileIntersection(file_name, deck_after)

   // Get only the cards that were created in the file.
   deck_after.OuterLeftJoin(deck_before)

   // Change the insert index based on if the current card was removed or not.
   card_index := 0
   if cur_hash == d.TopHash() {
      card_index = 1
   }

   // Check if the current card was removed or not.
   for i := deck_after.Len() - 1; i >= 0; i-- {
      new_card := deck_after.GetCard(i)
      d.InsertCard(new_card, card_index)
      d.AddMetaIfNil(new_card.HashStr(), cur_meta)
   }

   return nil
}
