package core

import "testing"
import "strings"
import "fmt"

const fullSum = "c435597dd9718c64b135087e944fd614716219d2f6663c1596e0d8f9c95e5681"
const halfSum = "c435597dd9718c64b135087e944fd614"

var f1 = "hello there | i'm a beard"
var f2 = "hello"
var f3 = "i'm um | hello"
var f4 = "alan the great | sy shoe yu"
var f5 = "a | b | c | d e | f | | g"

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
	c, err := NewCards("", f1)
	if err != nil {
		t.FailNow()
	}

	txt := c[0].String()
	if txt != "hello there | i'm a beard" {
		t.Fail()
	}

	cardSum := fmt.Sprintf("%x", c[0].Hash())

	if fullSum != cardSum {
		t.Fail()
	}
	if c[0].HashStr() != halfSum {
		t.Fail()
	}
	if !c[0].HasAnswer() {
		t.Fail()
	}
}

func TestDeck(t *testing.T) {
	d := NewDeck()
	d.AddNewCards("afile", f1)
	if d.GetCard(0).GetFactRaw(0) != "hello there" {
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
	d.AddNewCards("nofile", f1)
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
	d.AddNewCards("afile", f1)
	d.AddNewCards("afile", f2)
	d.AddNewCards("afile", f3)
	d.AddNewCards("afile", f4)
	d.TopToEnd()
	if d.TopCard().GetFactRaw(0) != "hello" {
		panic("Bad moves")
	}
	if d.Len() != 4 {
		panic("Bad len")
	}
	d.TopTo(10)
	if d.TopCard().GetFactRaw(0) != "i'm um" {
		panic("Bad moves")
	}
	d.TopTo(1)
	if d.TopCard().GetFactRaw(0) != "alan the great" {
		panic("Bad moves")
	}
	d.TopTo(1)
	if d.Len() != 4 {
		panic("Bad len")
	}
	if d.TopCard().GetFactRaw(0) != "i'm um" {
		panic("Bad moves")
	}
	d.TopToEnd()
	d.TopToEnd()
	if d.TopCard().GetFactRaw(0) != "hello there" {
		panic("Bad moves")
	}
}

func TestInsertCard(t *testing.T) {
	d := NewDeck()
	if c, err := NewCards("", f5); err != nil {
      panic("Should not error new card.")
   } else {
      if e := d.InsertCard(c[0], 5); e != nil    { panic("Should have inserted.") }
      if d.TopCard().HashStr() != c[0].HashStr() { panic("Card not inserted.") }
   }

	if c, err := NewCards("", f3); err != nil {
      panic("Should not error new card.")
   } else {
      if e := d.InsertCard(c[0], 0); e != nil    { panic("Should have inserted.") }
      if d.TopCard().HashStr() != c[0].HashStr() { panic("Card not inserted.") }
   }

	if c, err := NewCards("", f2); err != nil {
      panic("Should not error new card.")
   } else {
      if e := d.InsertCard(c[0], -100); e != nil    { panic("Should have inserted.") }
      if d.TopCard().HashStr() != c[0].HashStr() { panic("Card not inserted.") }
   }
}

func TestDoubleInsertCard(t *testing.T) {
	d := NewDeck()
	c1, _ := NewCards("file1", f3)
	c2, _ := NewCards("file2", f3)
   d.InsertCard(c1[0], 10)
   if err := d.InsertCard(c2[0], -10); err == nil || d.Len() != 1 {
      panic("Same card should have not been inserted twice!")
   }

   if err := d.AddNewCards("file3", f3); err == nil || d.Len() != 1 {
      panic("Card already exists and should not have been added.")
   }
}

func createDeck() *Deck {
	d := NewDeck()
	d.AddNewCards("a", f1)
	d.AddNewCards("b", f2)
	d.AddNewCards("a", f3)
	d.AddNewCards("c", f4)
	d.AddNewCards("b", f5)
   return d
}

func TestSwap(t *testing.T) {
	d := createDeck()
   if d.GetCard(0).String() != f1 && d.GetCard(1).String() != f2 {
      panic("Create deck doesn't have expected values.")
   }

   d.Swap(0,1)

   if d.GetCard(0).String() != f2 && d.GetCard(1).String() != f1 {
      panic("Swap didn't work.")
   }
}

func TestClone(t *testing.T) {
   d1 := createDeck()
   d2 := NewDeck()
   d2.Clone(d1)

   if d1.GetHash(0) != d2.GetHash(0) || d1.GetHash(1) != d2.GetHash(1) {
      panic("Cloned deck should be equal to original.")
   }
}

func TestCopy(t *testing.T) {
   d1 := createDeck()
   d2 := d1.Copy()

   if d1.GetHash(0) != d2.GetHash(0) || d1.GetHash(1) != d2.GetHash(1) {
      panic("Cloned deck should be equal to original.")
   }
}

func TestFilterNumber(t *testing.T) {
   d := createDeck()
   d.FilterNumber(1)
   if d.Len() != 1 {
      panic("Should have filtered down to one.")
   }
}

func TestFilterMemorize(t *testing.T) {
   d := createDeck()
   d.FilterOutMemorize()
   if d.Len() != 0 {
      panic("All were memorize.")
   }
}

func TestTop(t *testing.T) {
   d := createDeck()
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
   d := createDeck()
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
   d := createDeck()
	a := NewDefaultMeta("sm2")
   a.Next = a.Next.AddDate(1,0,0)
   d.AddMetaIfNil(d.TopHash(), a)

   oldLen := d.Len()
   d.FilterOutDone()
   newLen := d.Len()

   if oldLen - newLen != 1 {
      panic("There should have been one card done.")
   }

   if err := d.Forget(4); err == nil {
      panic("There is no 4 index. The highest is 3.")
   }

   if err := d.Forget(3); err != nil {
      panic("3 should have been a valid index.")
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

func TestEsc(t *testing.T) {
   raw := "in c, what is 1 \\| 2 | 3"
   esc := "in c, what is 1 | 2 | 3"
   rawFirst := "in c, what is 1 \\| 2"
   escFirst := "in c, what is 1 | 2"

	c, _ := NewCards("file1", raw)
   if c[0].GetFactEsc(0) != escFirst {
      panic("Fact not prettified/escaped.")
   }

   if c[0].GetFactRaw(0) != rawFirst {
      panic("Raw fact not what it was originally.")
   }

   if c[0].GetFactRaw(-1) != "" || c[0].GetFactEsc(-1) != "" || c[0].GetFactRaw(2) != "" || c[0].GetFactEsc(2) != "" {
      panic("Out of bounds esc and raw facts didn't work.")
   }

	if strings.Join(c[0].GetFactsRaw(), " | ") != raw {
      panic("Raw facts not preserved.")
   }

	if strings.Join(c[0].GetFactsEsc(), " | ") != esc {
      panic("Esc facts not preserved.")
   }
}
