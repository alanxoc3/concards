package stack

import (
	"sort"
	"time"

	"github.com/alanxoc3/concards/internal"
)

type Stack struct {
	mainKey   key                   // Current time and index.
	nextIndex int                   // The next index to insert.
	review    []internal.Hash       // Hashes to review ordered by date.
	future    []internal.Hash       // Hashes you have reviewed. Ordered by date.
	mapper    map[internal.Hash]key // All cards in this session.
}

func NewStack(t time.Time) Stack {
	return Stack{newMainKey(t), 0, []internal.Hash{}, []internal.Hash{}, map[internal.Hash]key{}}
}

func (s *Stack) SetTime(t time.Time) {
	s.mainKey = newMainKey(t)
}

func (s Stack) clone(o Stack) {
	s.mainKey = o.mainKey.clone()

	s.review = make([]internal.Hash, len(o.review))
	for i, v := range o.review {
		s.review[i] = v
	}

	s.future = make([]internal.Hash, len(o.future))
	for i, v := range o.future {
		s.future[i] = v
	}

	s.mapper = map[internal.Hash]key{}
	for k, v := range o.mapper {
		s.mapper[k] = v.clone()
	}
}

func (s *Stack) insertKey(h internal.Hash, k key) {
	// Step 1: Set the map.
	s.mapper[h] = k

	// Step 2: Add to the right stack.
	if k.after(s.mainKey) {
		s.future = insertSorted(s.future, h, func(i int) bool {
			return s.mapper[s.future[i]].after(k)
		})
	} else {
		s.review = insertSorted(s.review, h, func(i int) bool {
			return s.mapper[s.review[i]].before(k)
		})
	}
}

// Requires current card to have an entry in the predict map.
func (s *Stack) Insert(h internal.Hash, t time.Time) {
	if _, exist := s.mapper[h]; !exist {
		s.insertKey(h, key{t, s.nextIndex})
		s.nextIndex += 1 // TODO: Add max, this could technically overflow.
	}
}

// Guarantees a list in the insertion order.
func (s *Stack) List() []internal.Hash {
	hashes := []internal.Hash{}
	hashes = append(hashes, s.review...)
	hashes = append(hashes, s.future...)

	sort.Slice(hashes, func(i, j int) bool {
		return s.mapper[hashes[i]].index < s.mapper[hashes[j]].index
	})

	return hashes
}

func removeHashFromSlice(slice []internal.Hash, h internal.Hash) []internal.Hash {
	for i, v := range slice {
		if v == h {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func (s *Stack) Update(h internal.Hash, t time.Time) {
	if k, exist := s.mapper[h]; exist {
		// Step 1: Delete from slices.
		s.future = removeHashFromSlice(s.future, h)
		s.review = removeHashFromSlice(s.review, h)

		// Step 2: Re-insert into key map & lists, preserving insertion index.
		s.insertKey(h, key{t, k.index})
	}
}

// TODO do me pleeez.
func (s *Stack) FutureToReview() {
	if k, exist := s.mapper[h]; exist {
		// Step 1: Delete from slices.
		s.future = removeHashFromSlice(s.future, h)
		s.review = removeHashFromSlice(s.review, h)

		// Step 2: Re-insert into key map & lists, preserving insertion index.
		s.insertKey(h, key{t, k.index})
	}
}

func insertSorted(hs []internal.Hash, h internal.Hash, lessFunc func(int) bool) []internal.Hash {
	i := sort.Search(len(hs), lessFunc)
	hs = append(hs, internal.Hash{})
	copy(hs[i+1:], hs[i:])
	hs[i] = h
	return hs
}
