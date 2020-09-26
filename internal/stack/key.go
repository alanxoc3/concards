package stack

import "time"

type key struct {
	time  time.Time
	index int
}

func (k key) before(o key) bool {
	return k.time.Before(o.time) || k.time.Equal(o.time) && k.index < o.index
}

func (k key) after(o key) bool {
	return k.time.After(o.time) || k.time.Equal(o.time) && k.index > o.index
}

func (k key) clone() key {
	return key{k.time, k.index}
}

func newMainKey(t time.Time) key {
   // -1 is to put it before all cards.
	return key{t, -1}
}
