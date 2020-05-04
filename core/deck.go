package core

import "fmt"

func removeIndex(s []string, index int) []string {
    return append(s[:index], s[index+1:]...)
}

// refs contains checksums.
// Cmap maps checksums to cards.
// Mmap maps checksums to cards.
type Deck struct {
   refs  []string // This should be a byte array instead of a string!
   Cmap  map[string]*Card
   Mmap  map[string]*Meta
}

func NewDeck() *Deck {
   return &Deck{
      Cmap: map[string]*Card{},
      Mmap: map[string]*Meta{},
      refs: []string{},
   }
}

func (d *Deck) Forget(i int) error {
   if i >= 0 && i < len(d.refs) {
      delete(d.Mmap, d.refs[i])
      return nil
   } else {
      return fmt.Errorf("Can't forget. Index is out of bounds.")
   }
}

func (d *Deck) DelCard(i int) error {
   if i >= 0 && i < len(d.refs) {
      delete(d.Cmap, d.refs[i])
      d.refs = removeIndex(d.refs, i)
      return nil
   } else {
      return fmt.Errorf("Can't delete. Index is out of bounds.")
   }
}

func (d *Deck) AddCard(c *Card) error {
   hash := c.HashStr()
   _, exists := d.Cmap[hash]
   if !exists {
      d.Cmap[hash] = c
      d.refs = append(d.refs, hash)
      return nil
   } else {
      return fmt.Errorf("Card already exists in deck")
   }
}

func (d *Deck) AddFacts(facts [][]string, file string) error {
   if c, err := NewCard(facts, file); err == nil {
      return d.AddCard(c)
   }
   return nil
}

func (d *Deck) AddMeta(h string, m *Meta) {
   d.Mmap[h] = m
}

func (d *Deck) Len() int {
   return len(d.refs)
}

func (d *Deck) Swap(i, j int) {
   d.refs[i], d.refs[j] = d.refs[j], d.refs[i]
}

func (d *Deck) Get(i int) (h string, c *Card, m *Meta) {
   if i >= 0 && i < d.Len() {
      h = d.refs[i]
      c = d.Cmap[h]
      m = d.Mmap[h]
   }
   return
}

func (d *Deck) GetHash(i int) (h string) {
   h, _, _ = d.Get(i)
   return
}

func (d *Deck) GetCard(i int) (c *Card) {
   _, c, _ = d.Get(i)
   return
}

func (d *Deck) GetMeta(i int) (m *Meta) {
   _, _, m = d.Get(i)
   return
}

func (d *Deck) Copy() *Deck {
   n := NewDeck()
   for _, v := range d.refs { n.refs = append(n.refs, v) }
   for k, v := range d.Cmap { n.Cmap[k] = v }
   for k, v := range d.Mmap { n.Mmap[k] = v }
   return n
}

func (d *Deck) Clone(o *Deck) {
   d.refs = []string{}
   d.Cmap = map[string]*Card{}
   d.Mmap = map[string]*Meta{}

   for _, v := range o.refs { d.refs = append(d.refs, v) }
   for k, v := range o.Cmap { d.Cmap[k] = v }
   for k, v := range o.Mmap { d.Mmap[k] = v }
}
