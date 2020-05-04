package core

import "time"

// TODO: Refactor below functions to be cooler.
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
      if d.GetCard(i).File == path {
         refs = append(refs, i)
      }
   }

   d.delIndexes(refs)
}

func (d *Deck) FilterOutMemorize() {
   refs := []int{}

   for i :=0; i < d.Len(); i++ {
      if d.GetMeta(i) == nil {
         refs = append(refs, i)
      }
   }

   d.delIndexes(refs)
}

func (d *Deck) FilterOutReview() {
   refs := []int{}

   for i :=0; i < d.Len(); i++ {
      m := d.GetMeta(i)
      if m != nil && m.Next.Before(time.Now()) {
         refs = append(refs, i)
      }
   }

   d.delIndexes(refs)
}

func (d *Deck) FilterOutDone() {
   refs := []int{}

   for i :=0; i < d.Len(); i++ {
      m := d.GetMeta(i)
      if m != nil && m.Next.After(time.Now()) {
         refs = append(refs, i)
      }
   }

   d.delIndexes(refs)
}
