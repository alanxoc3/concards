package deck

import (
	"time"

	"github.com/alanxoc3/concards/internal"
)

type stack struct {
	review []internal.Hash // Hashes to review ordered by date.
	future []internal.Hash // Hashes you have reviewed. Ordered by date.
}

func (s stack) clone(o stack) {
	s.review = make([]internal.Hash, len(o.review))
	for i, v := range o.review {
		s.review[i] = v
	}

	s.future = make([]internal.Hash, len(o.future))
	for i, v := range o.future {
		s.future[i] = v
	}
}

func (s *stack) insertIntoReview(h internal.Hash, m predictMap) {
	next := m[h].Next()
	s.review = insertSorted(s.review, h, func(i int) bool {
		return m[s.review[i]].Next().Before(next)
	})
}

func (s *stack) insertIntoFuture(h internal.Hash, m predictMap) {
	next := m[h].Next()
	s.future = insertSorted(s.future, h, func(i int) bool {
		return m[s.future[i]].Next().After(next)
	})
}

// Requires current card to have an entry in the predict map.
func (s *stack) insert(h internal.Hash, m predictMap, now time.Time) {
	if beforeOrEqual(m[h].Next(), now) {
		s.insertIntoReview(h, m)
	} else {
		s.insertIntoFuture(h, m)
	}
}

func (s *stack) hashList() []internal.Hash {
	hashes := []internal.Hash{}
	hashes = append(hashes, s.review...)
	hashes = append(hashes, s.future...)
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

func (s *stack) refreshHash(h internal.Hash, m predictMap, now time.Time) {
   s.future = removeHashFromSlice(s.future, h)
   s.review = removeHashFromSlice(s.review, h)
   s.insert(h, m, now)
}

/*
func (s *stack) sort(m predictMap) []internal.Hash {
   sort.SliceStable(s.review, func(i, j, int) {
		return m[s.review[i]].Next().Before(m[s.review[j]].Next())
   })

   sort.SliceStable(s.future, func(i, j, int) {
		return m[s.future[i]].Next().After(m[s.future[j]].Next())
   })
}
*/
