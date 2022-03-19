package stack

import (
	"sort"
	"time"

	"github.com/alanxoc3/concards/internal"
)

type Stack struct {
	now       time.Time                   // Current time.
	review    []internal.Hash             // Hashes to review ordered by date.
	memorize  []internal.Hash             // Brand new hashes.
	future    []internal.Hash             // Hashes you have reviewed ordered by date.
	mapper    map[internal.Hash]time.Time // All cards in this session.
}

func NewStack(t time.Time) Stack {
	return Stack{t, []internal.Hash{}, []internal.Hash{}, []internal.Hash{}, map[internal.Hash]time.Time{}}
}

// A stack is sensitive to the time, so changing the time requires all the internals to be kept in sync.
func (s *Stack) SetTime(t time.Time) {
    if s.now.Before(t) {
        s.now = t
		for len(s.future) > 0 && (s.mapper[s.future[0]].Before(t) || s.mapper[s.future[0]].Equal(t)) {
			hash := s.future[0]
			s.future = s.future[1:]
			s.insertHash(hash, s.mapper[hash])
		}
    } else if s.now.After(t) {
        s.now = t
		for len(s.review) > 0 && s.mapper[s.review[0]].After(t) {
			hash := s.review[0]
			s.review = s.review[1:]
			s.insertHash(hash, s.mapper[hash])
		}
    }
}

func (s *Stack) Time() time.Time {
	return s.now
}

func (s *Stack) Top() *internal.Hash {
    if len(s.memorize) > 0 {
        h := s.memorize[0]
        return &h
    } else if len(s.review) > 0 {
        h := s.review[0]
        return &h
    } else {
        return nil
    }
}

func (s *Stack) Pop() *internal.Hash {
    if t := s.Top(); t != nil {
        s.memorize = removeHashFromSlice(s.memorize, *t)
        s.review = removeHashFromSlice(s.review, *t)
        return t
    }

    return nil
}

func (s *Stack) Clone(o Stack) {
	s.now = o.now

	s.memorize = make([]internal.Hash, len(o.memorize))
	for i, v := range o.memorize {
		s.memorize[i] = v
	}

	s.review = make([]internal.Hash, len(o.review))
	for i, v := range o.review {
		s.review[i] = v
	}

	s.future = make([]internal.Hash, len(o.future))
	for i, v := range o.future {
		s.future[i] = v
	}

	s.mapper = map[internal.Hash]time.Time{}
	for k, v := range o.mapper {
		s.mapper[k] = v
	}
}

// Guarantees a list sorted by date to be reviewed. Unreviewed cards are considered to be reviewed at current time, so they are in between reviewed and future hashes.
func (s *Stack) List() []internal.Hash {
	hashes := make([]internal.Hash, 0, s.Capacity())

	for i := len(s.review)-1; i >= 0; i-- {
        hashes = append(hashes, s.review[i])
	}

	hashes = append(hashes, s.memorize...)
	hashes = append(hashes, s.future...)

	return hashes
}

// Updates both the current time that the stack holds, and also the card that was reviewed.
func (s *Stack) Upsert(h internal.Hash, newTime time.Time) internal.Hash {
	if oldTime, exist := s.mapper[h]; exist && !oldTime.Equal(newTime) {
		s.memorize = removeHashFromSlice(s.memorize, h)
		s.review   = removeHashFromSlice(s.review,   h)
		s.future   = removeHashFromSlice(s.future,   h)
		s.insertHash(h, newTime)
	} else if !exist {
		s.insertHash(h, newTime)
	}

	return h
}

// The total number of cards loaded.
func (s *Stack) Capacity() int { return len(s.future) + len(s.review) + len(s.memorize) }

// The number of cards left to review.
func (s *Stack) Len() int      { return len(s.review) + len(s.memorize) }

// PRIVATE STUFF BELOW
type stackType int8
const (
	cNone stackType = iota
	cMemorize
	cReview
	cFuture
)

// inserts into the correct stack.
func (s *Stack) insertHash(h internal.Hash, t time.Time) {
	s.mapper[h] = t

	if t.IsZero() {
		s.memorize = append(s.memorize, h)
	} else if t.After(s.now) || t.Equal(s.now) {
		s.future = insertSorted(s.future, h, func(i int) bool {
        	return s.mapper[s.future[i]].After(t)
		})
	} else {
		s.review = insertSorted(s.review, h, func(i int) bool {
        	return s.mapper[s.review[i]].Before(t)
		})
	}
}

func removeHashFromSlice(slice []internal.Hash, h internal.Hash) []internal.Hash {
	for i, v := range slice {
		if v == h {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// inserts at the earliest spot less func returns true.
func insertSorted(hs []internal.Hash, h internal.Hash, lessFunc func(int) bool) []internal.Hash {
	i := sort.Search(len(hs), lessFunc)
	hs = append(hs, internal.Hash{})
	copy(hs[i+1:], hs[i:])
	hs[i] = h
	return hs
}
