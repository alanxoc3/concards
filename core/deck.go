package core

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

func (d *Deck) Forget(i int) {
   delete(d.Mmap, d.refs[i])
}

func (d *Deck) DelCard(i int) {
   d.refs = removeIndex(d.refs, i)
}

func (d *Deck) AddCard(c *Card) {
   hash := c.HashStr()
   d.Cmap[hash] = c
   d.refs = append(d.refs, hash)
}

func (d *Deck) AddFacts(facts [][]string, file string) {
   if c, err := NewCard(facts, file); err == nil {
      d.AddCard(c)
   }
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
