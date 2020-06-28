package core

import "testing"
import "strings"
import "fmt"

const fullSum = "c6cd355e32654cb4ba506b529ff32288971420ead2e36fdc69e802e9e7510315"
const halfSum = "c6cd355e32654cb4ba506b529ff32288"

var f1 = "hello there @ i'm a beard"
var f2 = "hello"
var f3 = "i'm um @ hello"
var f4 = "alan the great @ sy shoe yu"
var f5 = "a @ b @ c @ d e @ f @ @ g"

func TestMeta(t *testing.T) {
	a := NewMeta("2020-01-01T00:00:00Z", "0", "sm2", []string{"2.5"})

	if a.Name != "sm2" {
		t.Fail()
	}
	if a.Streak != 0 {
		t.Fail()
	}
	if a.Params[0] != "2.5" {
		t.Fail()
	}
	if a.IsZero() {
		t.Fail()
	}
	if a.String() != "2020-01-01T00:00:00Z 0 sm2 2.5" {
		t.Fail()
	}
}

func TestParse(t *testing.T) {
	if intOrDefault("123", 99) != 123 {
		t.Fail()
	}
	if intOrDefault("123a", 99) != 99 {
		t.Fail()
	}
}

func TestCard(t *testing.T) {
	c, err := NewCard("", f1)
	if err != nil {
		t.FailNow()
	}

	txt := c.String()
	if txt != "hello there @ i'm a beard" {
		t.Fail()
	}

	cardSum := fmt.Sprintf("%x", c.Hash())

	if fullSum != cardSum {
		t.Fail()
	}
	if c.HashStr() != halfSum {
		t.Fail()
	}
	if !c.HasAnswer() {
		t.Fail()
	}
}

func TestDeck(t *testing.T) {
	d := NewDeck()
	d.AddCardFromSides("afile", f1, false)
	if d.GetCard(0).GetQuestion() != "hello there" {
		t.Fail()
	}
	if !d.GetCard(0).HasAnswer() {
		t.Fail()
	}
	if d.GetMeta(0) != nil {
		t.Fail()
	}

	testmeta := NewMeta("", "", "sm2", []string{})
	d.AddMeta(d.GetHash(0), testmeta)
	if d.GetMeta(0).IsZero() {
		t.Fail()
	}
	md := strings.Fields(d.GetMeta(0).String())
	if len(md) != 3 {
		t.Fail()
	}
	if md[1] != "0" {
		t.Fail()
	}
	if md[2] != "sm2" {
		t.Fail()
	}

	d.ForgetTop()
	if d.GetMeta(0) != nil {
		t.Fail()
	}
	if !d.GetCard(0).HasAnswer() {
		t.Fail()
	}
	d.AddCardFromSides("nofile", f1, false)
	if d.Len() != 1 {
		t.Fail()
	}
	if d.GetCard(0).GetFile() != "afile" {
		t.Fail()
	}
	d.FilterOutFile("nofile")
	if d.Len() != 1 {
		t.Fail()
	}
	d.FilterOutFile("afile")
	if !d.IsEmpty() {
		t.Fail()
	}
	d.Del(0)
	if d.GetCard(0) != nil {
		t.Fail()
	}
}

func TestDeckMove(t *testing.T) {
	d := NewDeck()
	d.AddCardFromSides("afile", f1, false)
	d.AddCardFromSides("afile", f2, false)
	d.AddCardFromSides("afile", f3, false)
	d.AddCardFromSides("afile", f4, false)
	d.TopToEnd()
	if d.TopCard().GetQuestion() != "hello" {
		panic("Bad moves")
	}
	if d.Len() != 4 {
		panic("Bad len")
	}
	d.TopTo(10)
	if d.TopCard().GetQuestion() != "i'm um" {
		panic("Bad moves")
	}
	d.TopTo(1)
	if d.TopCard().GetQuestion() != "alan the great" {
		panic("Bad moves")
	}
	d.TopTo(1)
	if d.Len() != 4 {
		panic("Bad len")
	}
	if d.TopCard().GetQuestion() != "i'm um" {
		panic("Bad moves")
	}
	d.TopToEnd()
	d.TopToEnd()
	if d.TopCard().GetQuestion() != "hello there" {
		panic("Bad moves")
	}
}

func TestAddSubCards(t *testing.T) {
	d := NewDeck()
	d.AddCardFromSides("dat-file", f5, true)

	if d.TopCard().GetQuestion() != "a" {
		panic("Subcards were inserted before the parent card.")
	}
	if d.Len() != 6 {
		panic("Wrong number of sub cards inserted.")
	}
	d.DelTop()
	if d.TopCard().GetQuestion() != "b" {
		panic("Second card should be the first sub card.")
	}
	if d.TopCard().GetFact(1) != "a" {
		panic("Answer isn't the parent card.")
	}
	if d.Len() != 5 {
		panic("Delete didn't work.")
	}

	if d.GetCard(1).GetQuestion() != "c" {
		panic("Sub cards not inserted in the correct order.")
	}
	if d.GetCard(1).GetFact(1) != "a" {
		panic("Sub card doesn't have parent as the answer.")
	}

	if d.GetCard(2).GetQuestion() != "d e" {
		panic("Sub cards not inserted in the correct order.")
	}
	if d.GetCard(2).GetFact(1) != "a" {
		panic("Sub card doesn't have parent as the answer.")
	}

	if d.GetCard(3).GetQuestion() != "f" {
		panic("Sub cards not inserted in the correct order.")
	}
	if d.GetCard(3).GetFact(1) != "a" {
		panic("Sub card doesn't have parent as the answer.")
	}

	if d.GetCard(4).GetQuestion() != "g" {
		panic("Sub cards not inserted in the correct order.")
	}
	if d.GetCard(4).GetFact(1) != "a" {
		panic("Sub card doesn't have parent as the answer.")
	}
}

func TestInsertCard(t *testing.T) {
	d := NewDeck()
	if c, err := NewCard("", f5); err != nil {
      panic("Should not error new card.")
   } else {
      if e := d.InsertCard(c, 5); e != nil    { panic("Should have inserted.") }
      if d.TopCard().HashStr() != c.HashStr() { panic("Card not inserted.") }
   }

	if c, err := NewCard("", f3); err != nil {
      panic("Should not error new card.")
   } else {
      if e := d.InsertCard(c, 0); e != nil    { panic("Should have inserted.") }
      if d.TopCard().HashStr() != c.HashStr() { panic("Card not inserted.") }
   }

	if c, err := NewCard("", f2); err != nil {
      panic("Should not error new card.")
   } else {
      if e := d.InsertCard(c, -100); e != nil    { panic("Should have inserted.") }
      if d.TopCard().HashStr() != c.HashStr() { panic("Card not inserted.") }
   }
}

func TestDoubleInsertCard(t *testing.T) {
	d := NewDeck()
	c1, _ := NewCard("file1", f3)
	c2, _ := NewCard("file2", f3)
   d.InsertCard(c1, 10)
   if err := d.InsertCard(c2, -10); err == nil || d.Len() != 1 {
      panic("Same card should have not been inserted twice!")
   }

   if errs := d.AddCardFromSides("file3", f3, false); len(errs) != 1 {
      panic("Card already exists and should not have been added.")
   }
}

func createDeck(includeSubcards bool) *Deck {
	d := NewDeck()
	d.AddCardFromSides("a", f1, includeSubcards)
	d.AddCardFromSides("b", f2, includeSubcards)
	d.AddCardFromSides("a", f3, includeSubcards)
	d.AddCardFromSides("c", f4, includeSubcards)
	d.AddCardFromSides("b", f5, includeSubcards)
   return d
}

func TestSwap(t *testing.T) {
	d := createDeck(false)
   if d.GetCard(0).String() != f1 && d.GetCard(1).String() != f2 {
      panic("Create deck doesn't have expected values.")
   }

   d.Swap(0,1)

   if d.GetCard(0).String() != f2 && d.GetCard(1).String() != f1 {
      panic("Swap didn't work.")
   }
}

func TestClone(t *testing.T) {
   d1 := createDeck(false)
   d2 := NewDeck()
   d2.Clone(d1)

   if d1.GetHash(0) != d2.GetHash(0) || d1.GetHash(1) != d2.GetHash(1) {
      panic("Cloned deck should be equal to original.")
   }
}

func TestCopy(t *testing.T) {
   d1 := createDeck(false)
   d2 := d1.Copy()

   if d1.GetHash(0) != d2.GetHash(0) || d1.GetHash(1) != d2.GetHash(1) {
      panic("Cloned deck should be equal to original.")
   }
}

func TestFilterNumber(t *testing.T) {
   d := createDeck(true)
   d.FilterNumber(1)
   if d.Len() != 1 {
      panic("Should have filtered down to one.")
   }
}

func TestFilterMemorize(t *testing.T) {
   d := createDeck(true)
   d.FilterOutMemorize()
   if d.Len() != 0 {
      panic("All were memorize.")
   }
}

func TestTop(t *testing.T) {
   d := createDeck(false)
	m1 := NewMeta("2020-01-01T00:00:00Z", "0", "sm2", []string{"2.5"})
   d.AddMeta(d.TopHash(), m1)

   h, c, m2 := d.Top()

   if h != d.TopHash() {
      panic("Top returned bad hash.")
   }

   if c.String() != f1 && c != d.TopCard() {
      panic("Top returned bad card.")
   }

   if m1 != m2 || m2.Next != d.TopMeta().Next {
      panic("Top returned bad meta.")
   }
}

func TestFilterReview(t *testing.T) {
   d := createDeck(true)
	a := NewMeta("2020-01-01T00:00:00Z", "0", "sm2", []string{"2.5"})
   d.AddMeta(d.TopHash(), a)

   oldLen := d.Len()
   d.FilterOutReview()
   newLen := d.Len()

   if oldLen - newLen != 1 {
      panic("There should have been one card to review.")
   }
}

func TestFilterDone(t *testing.T) {
   d := createDeck(true)
	a := NewDefaultMeta("sm2")
   a.Next = a.Next.AddDate(1,0,0)
   d.AddMetaIfNil(d.TopHash(), a)

   oldLen := d.Len()
   d.FilterOutDone()
   newLen := d.Len()

   if oldLen - newLen != 1 {
      panic("There should have been one card done.")
   }

   if err := d.Forget(12); err == nil {
      panic("There is no 12 index. The highest is 11.")
   }

   if err := d.Forget(11); err != nil {
      panic("11 should have been a valid index.")
   }
}

func TestSm2Exec(t *testing.T) {
   // Testing NO.
	a := NewMeta("2020-01-01T00:00:00Z", "1", "sm2", []string{"2.5"})
   a, _ = a.Exec(NO)
   if a.Params[0] != "1.96" {
      panic("Sm2 returned the wrong weight.")
   }

   // Testing NO, streak go down.
	a = NewMeta("2020-01-01T00:00:00Z", "-3", "sm2", []string{"2.5"})
   a, _ = a.Exec(NO)
   if a.Params[0] != "1.96" && a.Streak != -4 {
      panic("Sm2 returned the wrong weight.")
   }

   // Testing IDK.
	a = NewMeta("2020-01-01T00:00:00Z", "1", "sm2", []string{"2.5"})
   a, _ = a.Exec(IDK)
   if a.Params[0] != "2.36" {
      panic("Sm2 returned the wrong weight.")
   }

   // Testing YES.
	a = NewMeta("2020-01-01T00:00:00Z", "3", "sm2", []string{"2.5"})
   a, _ = a.Exec(YES)
   if a.Params[0] != "2.60" {
      panic("Sm2 returned the wrong weight.")
   }

   // Testing YES, negative streak.
	a = NewMeta("2020-01-01T00:00:00Z", "-1", "sm2", []string{"2.5"})
   a, _ = a.Exec(YES)
   if a.Params[0] != "2.50" {
      panic("Sm2 returned the wrong weight.")
   }

   // Should not go lower than the lowest weight.
	a = NewMeta("2020-01-01T00:00:00Z", "1", "sm2", []string{"1.3"})
   a, _ = a.Exec(NO)
   if a.Params[0] != "1.30" {
      panic("Sm2 returned the wrong weight.")
   }
}
