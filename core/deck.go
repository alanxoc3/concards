package core

import "fmt"

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

// refs contains checksums.
// Cmap maps checksums to cards.
// Mmap maps checksums to cards.
type Deck struct {
	refs []string // This should be a byte array instead of a string!
	Cmap map[string]*Card
	Mmap map[string]*Meta
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

func (d *Deck) Del(i int) error {
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

func (d *Deck) InsertCard(c *Card, i int) error {
	hash := c.HashStr()

   if i < 0 {
      i = 0
   }

	_, exists := d.Cmap[hash]
	if !exists {
		d.Cmap[hash] = c

      if i >= d.Len() {
         d.refs = append(d.refs, hash)
      } else {
         d.refs = append(d.refs, "")
         copy(d.refs[i+1:], d.refs[i:])
         d.refs[i] = hash
      }

		return nil
	} else {
		return fmt.Errorf("Card already exists in deck")
	}
}

func (d *Deck) AddCardFromSides(file string, sides string) []error {
	errors := []error{}
	if c, createErr := NewCard(file, sides); createErr == nil {
		cards := []*Card{c}

		for _, c := range cards {
			if addErr := d.AddCard(c); addErr != nil {
				errors = append(errors, addErr)
			}
		}
	} else {
		errors = append(errors, createErr)
	}
	return errors
}

func (d *Deck) AddMeta(h string, m *Meta) {
	d.Mmap[h] = m
}

func (d *Deck) AddMetaIfNil(h string, m *Meta) {
	if _, ok := d.Mmap[h]; !ok {
		d.AddMeta(h, m)
	}
}

func (d *Deck) Len() int {
	return len(d.refs)
}

func (d *Deck) IsEmpty() bool {
	return d.Len() == 0
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
	for _, v := range d.refs {
		n.refs = append(n.refs, v)
	}
	for k, v := range d.Cmap {
		n.Cmap[k] = v
	}
	for k, v := range d.Mmap {
		n.Mmap[k] = v
	}
	return n
}

func (d *Deck) Clone(o *Deck) {
	d.refs = []string{}
	d.Cmap = map[string]*Card{}
	d.Mmap = map[string]*Meta{}

	for _, v := range o.refs {
		d.refs = append(d.refs, v)
	}
	for k, v := range o.Cmap {
		d.Cmap[k] = v
	}
	for k, v := range o.Mmap {
		d.Mmap[k] = v
	}
}

// Top shortcuts
func (d *Deck) Top() (string, *Card, *Meta) { return d.Get(0) }
func (d *Deck) TopHash() string             { return d.GetHash(0) }
func (d *Deck) TopCard() *Card              { return d.GetCard(0) }
func (d *Deck) TopMeta() *Meta              { return d.GetMeta(0) }
func (d *Deck) DelTop() error               { return d.Del(0) }
func (d *Deck) ForgetTop() error            { return d.Forget(0) }

func (d *Deck) TopToEnd() {
	if len(d.refs) > 1 {
		d.refs = append(d.refs[1:], d.refs[0])
	}
}

func (d *Deck) TopTo(i int) {
	if l := len(d.refs); l > 1 && i > 0 {
		if i >= l {
			i = l - 1
		}
		v := d.refs[0]
		copy(d.refs, d.refs[1:i+1])
		d.refs[i] = v
	}
}
