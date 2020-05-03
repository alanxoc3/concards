package core

func removeIndex(s []string, index int) []string {
    return append(s[:index], s[index+1:]...)
}

// refs contains checksums.
// cmap maps checksums to cards.
// mmap maps checksums to cards.
type Deck struct {
   cmap  map[string]*Card // This should be a pointer!
   mmap  map[string]*Meta // This should be a pointer!
   refs  []string // This should be a byte array instead of a string!
}

func NewDeck() *Deck {
   return &Deck{
      cmap: map[string]*Card{},
      mmap: map[string]*Meta{},
      refs: []string{},
   }
}

func (d *Deck) Forget(i int) {
   delete(d.mmap, d.refs[i])
}

func (d *Deck) DelCard(i int) {
   d.refs = removeIndex(d.refs, i)
}

func (d *Deck) AddCard(c *Card) {
   hash := c.HashStr()
   d.cmap[hash] = c
   d.refs = append(d.refs, hash)
}

func (d *Deck) AddFacts(facts [][]string) {
   if c, err := NewCard(facts); err == nil {
      d.AddCard(c)
   }
}

func (d *Deck) AddMeta(h string, m Meta) {
   d.mmap[h] = &m
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
      c = d.cmap[h]
      m = d.mmap[h]
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
