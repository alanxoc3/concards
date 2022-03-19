package stack_test

import (
	"testing"
	"time"

	"github.com/alanxoc3/concards/internal"
	"github.com/alanxoc3/concards/internal/stack"
	"github.com/stretchr/testify/assert"
)

var DATE_EMPTY time.Time = time.Time{}
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
	assert.Zero(t, s.Len())
	assert.Zero(t, s.Capacity())
	assert.Nil(t, s.Pop())
}

func TestInsert(t *testing.T) {
	s := stack.NewStack(DATE_1)
   h := internal.NewHash("")
   s.Upsert(h, DATE_1)
	assert.Len(t, s.List(), 1)
}

func TestInsertSameFuture(t *testing.T) {
    s := stack.NewStack(DATE_2)
    h := internal.NewHash("")
    s.Upsert(h, DATE_2)
    assert.Equal(t, 0, s.Len())
    assert.Equal(t, 1, s.Capacity())
}

func TestInsertDifferentFuture(t *testing.T) {
	s := stack.NewStack(DATE_1)
   h := internal.NewHash("")
   s.Upsert(h, DATE_2)
	assert.Equal(t, 0, s.Len())
	assert.Equal(t, 1, s.Capacity())
}

func TestInsertDifferentReview(t *testing.T) {
	s := stack.NewStack(DATE_2)
   h := internal.NewHash("")
   s.Upsert(h, DATE_1)
	assert.Equal(t, h, *s.Top())
	assert.Equal(t, 1, s.Len())
	assert.Equal(t, 1, s.Capacity())
}

func TestPop(t *testing.T) {
	s := stack.NewStack(DATE_2)
   s.Upsert(internal.NewHash("ff"), DATE_1)
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

	s.Upsert(a, DATE_1)
	s.Upsert(b, DATE_EMPTY)
	s.Upsert(c, DATE_EMPTY)
	s.Upsert(d, DATE_2)
	s.Upsert(e, DATE_EMPTY)
	s.Upsert(f, DATE_EMPTY)

	list := []internal.Hash{}
	for s.Top() != nil {
		list = append(list, *s.Top())
		s.Pop()
	}

	assert.Equal(t, []internal.Hash{b, c, e, f, d, a}, list)
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
   g := internal.NewHash("5")

   s.Upsert(a, DATE_4)
   s.Upsert(b, DATE_3)
   s.Upsert(c, DATE_3)
   s.Upsert(d, DATE_1)
   s.Upsert(e, DATE_2)
   s.Upsert(f, DATE_1)
   s.Upsert(g, DATE_4)

   assert.Equal(t, []internal.Hash{f, d, e, b, c, a, g}, s.List())
}

func TestUpdate(t *testing.T) {
	s := stack.NewStack(DATE_2)
   f := internal.NewHash("f")
   e := internal.NewHash("e")
   d := internal.NewHash("d")

   s.Upsert(f, DATE_2)
   s.Upsert(e, DATE_2)
   s.Upsert(d, DATE_2)

   s.Upsert(f, DATE_1)
   s.Upsert(d, DATE_1)

   assert.Equal(t, 2, s.Len())
   assert.Equal(t, 3, s.Capacity())
   assert.Equal(t, f, *s.Top())
}

func TestUpdateFutureToReview(t *testing.T) {
	s := stack.NewStack(time.Date(2020,1,1,0,0,0,0,time.UTC))
   f := internal.NewHash("f")
   e := internal.NewHash("e")
   d := internal.NewHash("d")

   s.Upsert(f, time.Date(2020,1,1,0,0,0,1,time.UTC))
   s.Upsert(e, time.Date(2019,1,1,0,0,0,0,time.UTC))
   s.Upsert(d, time.Date(2019,1,1,0,0,0,0,time.UTC))
   assert.Equal(t, 3, s.Capacity())
   assert.Equal(t, 2, s.Len())

   s.SetTime(time.Date(2020,1,1,0,0,0,2,time.UTC))
   assert.Equal(t, 3, s.Capacity())
   assert.Equal(t, 3, s.Len())
   assert.Equal(t, f, *s.Top())
}
