package stack_test

import (
	"testing"
	"time"

	"github.com/alanxoc3/concards/internal"
	"github.com/alanxoc3/concards/internal/stack"
	"github.com/stretchr/testify/assert"
)

var DATE_0 time.Time = time.Time{}
var DATE_1 time.Time = time.Date(1, 1, 1, 0, 0, 1, 0, time.UTC)
var DATE_2 time.Time = time.Date(1, 1, 1, 0, 0, 2, 0, time.UTC)
var DATE_3 time.Time = time.Date(1, 1, 1, 0, 0, 3, 0, time.UTC)
var DATE_4 time.Time = time.Date(1, 1, 1, 0, 0, 4, 0, time.UTC)
var DATE_5 time.Time = time.Date(1, 1, 1, 0, 0, 5, 0, time.UTC)

func TestTime(t *testing.T) {
	s := stack.NewStack(DATE_1)
	assert.Equal(t, DATE_1, s.Time())
}

func TestSetTime(t *testing.T) {
	s := stack.NewStack(DATE_1)
	s.SetTime(DATE_2)
	assert.Equal(t, DATE_2, s.Time())
}

func TestEmpty(t *testing.T) {
	s := stack.NewStack(DATE_1)
	assert.Nil(t, s.Top())
	assert.Empty(t, s.List())
	assert.Zero(t, s.ReviewLen())
	assert.Zero(t, s.FutureLen())
	assert.False(t, s.Update(internal.NewHash(""), DATE_2, false))
}

func TestInsert(t *testing.T) {
	s := stack.NewStack(DATE_1)
   h := internal.NewHash("")
   s.Insert(h, DATE_1, false)
	assert.Len(t, s.List(), 1)
}

func TestInsertSameFuture(t *testing.T) {
	s := stack.NewStack(DATE_2)
   h := internal.NewHash("")
   s.Insert(h, DATE_2, false)
	assert.Equal(t, 0, s.ReviewLen())
	assert.Equal(t, 1, s.FutureLen())
}

func TestInsertDifferentFuture(t *testing.T) {
	s := stack.NewStack(DATE_1)
   h := internal.NewHash("")
   s.Insert(h, DATE_2, false)
	assert.Equal(t, 0, s.ReviewLen())
	assert.Equal(t, 1, s.FutureLen())
}

func TestInsertDifferentReview(t *testing.T) {
	s := stack.NewStack(DATE_2)
   h := internal.NewHash("")
   s.Insert(h, DATE_1, false)
	assert.Equal(t, h, *s.Top())
	assert.Equal(t, 1, s.ReviewLen())
	assert.Equal(t, 0, s.FutureLen())
}

func TestPop(t *testing.T) {
	s := stack.NewStack(DATE_2)
   s.Insert(internal.NewHash("ff"), DATE_1, false)
   s.Pop()
	assert.Len(t, s.List(), 0)
}

func TestClone(t *testing.T) {
	s1 := stack.NewStack(DATE_1)
	s2 := stack.NewStack(DATE_2)
   s1.Clone(s2)

	assert.Equal(t, s2, s1)
}

func TestInsertMemorizePriority(t *testing.T) {
	s := stack.NewStack(DATE_3)

	a := internal.NewHash("a")
	b := internal.NewHash("b")
	c := internal.NewHash("c")
	d := internal.NewHash("d")
	e := internal.NewHash("e")
	f := internal.NewHash("f")

	s.Insert(a, DATE_1, false)
	s.Insert(b, DATE_1, true)
	s.Insert(c, DATE_2, true)
	s.Insert(d, DATE_2, false)
	s.Insert(e, DATE_1, true)
	s.Insert(f, DATE_4, true)

	list := []internal.Hash{}
	for s.Top() != nil {
		list = append(list, *s.Top())
		s.Pop()
	}

	assert.Equal(t, []internal.Hash{c, b, e, d, a}, list)
}

// Insertion order should be in order of date, then insertion order.
func TestListInsertionOrder(t *testing.T) {
	s := stack.NewStack(DATE_3)
   a := internal.NewHash("a")
   b := internal.NewHash("b")
   c := internal.NewHash("c")
   d := internal.NewHash("d")
   e := internal.NewHash("e")
   f := internal.NewHash("f")

   s.Insert(a, DATE_4, false)
   s.Insert(b, DATE_3, false)
   s.Insert(c, DATE_3, false)
   s.Insert(d, DATE_1, false)
   s.Insert(e, DATE_2, false)
   s.Insert(f, DATE_1, false)

   assert.Equal(t, []internal.Hash{d, f, e, b, c, a}, s.List())
}

func TestUpdate(t *testing.T) {
	s := stack.NewStack(DATE_2)
   f := internal.NewHash("f")
   e := internal.NewHash("e")
   d := internal.NewHash("d")

   s.Insert(f, DATE_2, false)
   s.Insert(e, DATE_2, false)
   s.Insert(d, DATE_2, false)

   s.Update(f, DATE_1, true)
   s.Update(d, DATE_1, false)

   assert.Equal(t, 2, s.ReviewLen())
   assert.Equal(t, 1, s.FutureLen())
   assert.Equal(t, f, *s.Top())
}

func TestUpdateFutureToReview(t *testing.T) {
	s := stack.NewStack(time.Date(2020,1,1,0,0,0,0,time.UTC))
   f := internal.NewHash("f")
   e := internal.NewHash("e")
   d := internal.NewHash("d")

   s.Insert(f, time.Date(2020,1,1,0,0,0,1,time.UTC), false)
   s.Insert(e, time.Date(2019,1,1,0,0,0,0,time.UTC), false)
   s.Insert(d, time.Date(2019,1,1,0,0,0,0,time.UTC), false)
   assert.Equal(t, 1, s.FutureLen())

   s.SetTime(time.Date(2020,1,1,0,0,0,2,time.UTC))
   s.Update(e, DATE_1, false)

   assert.Equal(t, 0, s.FutureLen())
   assert.Equal(t, f, *s.Top())
}
