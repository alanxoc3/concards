package con

import "testing"
import "fmt"

const fullSum = "c6cd355e32654cb4ba506b529ff32288971420ead2e36fdc69e802e9e7510315"
const halfSum = "c6cd355e32654cb4ba506b529ff32288"

var facts = [][]string{
   {"hello", "there"},
   {"i'm", "a", "beard"},
}

func TestAlg(t *testing.T) {
   a := NewAlg("2020-01-01T00:00:00Z", "0", "sm2", []string{"2.5"})

   if a.Name != "sm2" { t.Fail() }
   if a.Streak != 0 { t.Fail() }
   if a.Params[0] != "2.5" { t.Fail() }
}

func TestParse(t *testing.T) {
   if intOrDefault("123", 99) != 123 { t.Fail() }
   if intOrDefault("123a", 99) != 99 { t.Fail() }
}

func TestCard(t *testing.T) {
   c, err := NewCard(facts)
   if err != nil { t.FailNow() }

   txt := c.KeyText()
   if txt != "hello there @ i'm a beard" { t.Fail() }

   cardSum := fmt.Sprintf("%x", c.Hash())

   if fullSum != cardSum { t.Fail() }
   if c.HashStr() != halfSum { t.Fail() }
   if !c.HasAnswer() { t.Fail() }
}

func TestDeck(t *testing.T) {
   d := NewDeck().AddFacts(facts)
   if d.GetCard(0).GetQuestion() != "hello there" { t.Fail() }
   if !d.GetCard(0).HasAnswer() { t.Fail() }
}
