package core

func (d *Deck) delIndexes(refs []int) {
   for i := len(refs)-1; i >= 0; i-- {
      d.DelCard(i)
   }
}

// Basically truncates a deck.
func (d *Deck) FilterNumber(param int) {
   refs := []int{}

   for i :=0; i < d.Len(); i++ {
      if i >= param {
         refs = append(refs, i)
      }
   }

   d.delIndexes(refs)
}

func (d *Deck) FilterOutFile(path string) {
   refs := []int{}

   for i :=0; i < d.Len(); i++ {
      c := d.GetCard(i)
      if c.File == path {
         refs = append(refs, i)
      }
   }

   d.delIndexes(refs)
}
