package stack

import "time"

type key struct {
	time  time.Time
	index int
	memorize bool
}

// Keys with a higher date go to the top.
// If 2 keys have the same date, the one with a smaller index will be at the top.
func (k key) reviewLess(o key) bool {
    if k.memorize && !o.memorize {
        return false
    } else if o.memorize && !k.memorize {
        return true
    }

    return k.time.Before(o.time) || k.time.Equal(o.time) && k.index > o.index
}

// Keys with a lower date go to the top.
// If 2 keys have the same date, the one with a smaller index will be at the top.
func (k key) futureLess(o key) bool {
	return k.time.After(o.time) || k.time.Equal(o.time) && k.index > o.index
}

func (k key) beforeTime(o key) bool {
	return k.time.Before(o.time)
}

func (k key) clone() key {
	return key{k.time, k.index, k.memorize}
}

func newMainKey(t time.Time) key {
   // -1 is to put it before all cards.
	return key{t, -1, false}
}
