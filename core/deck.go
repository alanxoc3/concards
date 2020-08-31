package core

import "fmt"
import "sort"

func removeIndex(s []string, index int) []string {
   return append(s[:index], s[index+1:]...)
}

// refs contains checksums.
// refsMap maps checksums to cards.
// MetaMap maps checksums to cards.
type Deck struct {
   refs []string // Hashes to review.
   reviews []string // Hashes you have reviewed. Ordered by date.
   refsMap map[string]*Card // All cards in this session.
   metaHist []MetaHist // Meta history.
   MetaMap map[string]*MetaAlg // All metas.
}

func NewDeck() *Deck {
   return &Deck{
      refs: []string{},
      reviews: []string{},
      refsMap: map[string]*Card{},
      metaHist: []MetaHist{},
      MetaMap: map[string]*MetaAlg{},
   }
}

func (d *Deck) Forget(i int) error {
   if i >= 0 && i < len(d.refs) {
      delete(d.MetaMap, d.refs[i])
      return nil
   } else {
      return fmt.Errorf("Can't forget. Index is out of bounds.")
   }
}

// Deletes from the deck.
func (d *Deck) Drop(i int) error {
   if i >= 0 && i < len(d.refs) {
      delete(d.refsMap, d.refs[i])
      d.refs = removeIndex(d.refs, i)
      return nil
   } else {
      return fmt.Errorf("Can't delete. Index is out of bounds.")
   }
}

func (d *Deck) Delay(i int) error {
   if i >= 0 && i < len(d.refs) {
      d.reviews = append(d.reviews, d.refs[i])
      d.refs = removeIndex(d.refs, i)

      sort.Slice(d.reviews, func(i, j int) bool {
         return d.MetaMap[d.reviews[i]].Next.Before(d.MetaMap[d.reviews[j]].Next)
      })
      return nil
   } else {
      return fmt.Errorf("Can't delay. Index is out of bounds.")
   }
}

func (d *Deck) AddCard(c *Card) error {
   hash := c.HashStr()
   _, exists := d.refsMap[hash]
   if !exists {
      d.refsMap[hash] = c
      d.refs = append(d.refs, hash)
      return nil
   } else {
      return fmt.Errorf("Card already exists in deck")
   }
}

func (d *Deck) InsertCard(c *Card, i int) error {
   hash := c.HashStr()
   if i < 0 { i = 0 }

   _, exists := d.refsMap[hash]
   if !exists {
      d.refsMap[hash] = c

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

func (d *Deck) AddNewCards(file string, sides string) error {
   if cards, err := NewCards(file, sides); err != nil {
      return err
   } else {
      for _, c := range cards {
         if addErr := d.AddCard(c); addErr != nil {
            err = addErr
         }
      }
      return err
   }
}

func (d *Deck) AddMeta(h string, m *MetaAlg) {
   d.MetaMap[h] = m
}

func (d *Deck) AddMetaIfNil(h string, m *MetaAlg) {
   if _, ok := d.MetaMap[h]; !ok {
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

func (d *Deck) Get(i int) (h string, c *Card, m *MetaAlg) {
   if i >= 0 && i < d.Len() {
      h = d.refs[i]
      c = d.refsMap[h]
      m = d.MetaMap[h]
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

func (d *Deck) GetMeta(i int) (m *MetaAlg) {
   _, _, m = d.Get(i)
   return
}

func (d *Deck) Copy() *Deck {
   n := NewDeck()
   for _, v := range d.refs {
      n.refs = append(n.refs, v)
   }
   for k, v := range d.refsMap {
      n.refsMap[k] = v
   }
   for k, v := range d.MetaMap {
      n.MetaMap[k] = v
   }
   return n
}

func (d *Deck) Clone(o *Deck) {
   d.refs = []string{}
   d.refsMap = map[string]*Card{}
   d.MetaMap = map[string]*MetaAlg{}

   for _, v := range o.refs {
      d.refs = append(d.refs, v)
   }
   for k, v := range o.refsMap {
      d.refsMap[k] = v
   }
   for k, v := range o.MetaMap {
      d.MetaMap[k] = v
   }
}

// Top shortcuts
func (d *Deck) Top() (string, *Card, *MetaAlg) { return d.Get(0) }
func (d *Deck) TopHash() string                { return d.GetHash(0) }
func (d *Deck) TopCard() *Card                 { return d.GetCard(0) }
func (d *Deck) TopMeta() *MetaAlg              { return d.GetMeta(0) }
func (d *Deck) DropTop() error                 { return d.Drop(0) }
func (d *Deck) DelayTop() error                { return d.Delay(0) }
func (d *Deck) ForgetTop() error               { return d.Forget(0) }

func (d *Deck) TopMetaOrDefault(defaultAlg string) *MetaAlg {
   m := d.TopMeta()
   if m == nil {
      return NewDefaultMetaAlg(d.TopHash(), defaultAlg)
   }
   return m
}

func (d *Deck) ExecTop(input bool, defaultAlg string) (*MetaAlg, error) {
   h := d.TopHash()

   if ma, e := d.TopMetaOrDefault(defaultAlg).Exec(input); e != nil {
      d.DropTop()
      return nil, e
   } else {
      d.AddMeta(h, ma)
      d.DelayTop()
      return ma, nil
   }
}

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
