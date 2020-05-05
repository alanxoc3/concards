package core

import "testing"
import "strings"
import "fmt"

const fullSum = "c6cd355e32654cb4ba506b529ff32288971420ead2e36fdc69e802e9e7510315"
const halfSum = "c6cd355e32654cb4ba506b529ff32288"

var f1 = [][]string{
   {"hello", "there"},
   {"i'm", "a", "beard"},
}

var f2 = [][]string{
   {"hello"},
}

var f3 = [][]string{
   {"i'm", "um"},
   {"hello"},
}

var f4 = [][]string{
   {"alan", "the", "great"},
   {"sy", "shoe", "yu"},
}

func TestMeta(t *testing.T) {
   a := NewMeta("2020-01-01T00:00:00Z", "0", "sm2", []string{"2.5"})

   if a.Name != "sm2" { t.Fail() }
   if a.Streak != 0 { t.Fail() }
   if a.Params[0] != "2.5" { t.Fail() }
   if a.IsZero() { t.Fail() }
   if a.String() != "2020-01-01T00:00:00Z 0 sm2 2.5" { t.Fail() }
}

func TestParse(t *testing.T) {
   if intOrDefault("123", 99) != 123 { t.Fail() }
   if intOrDefault("123a", 99) != 99 { t.Fail() }
}

func TestCard(t *testing.T) {
   c, err := NewCard(f1, "")
   if err != nil { t.FailNow() }

   txt := c.String()
   if txt != "hello there @ i'm a beard" { t.Fail() }

   cardSum := fmt.Sprintf("%x", c.Hash())

   if fullSum != cardSum { t.Fail() }
   if c.HashStr() != halfSum { t.Fail() }
   if !c.HasAnswer() { t.Fail() }
}

func TestDeck(t *testing.T) {
   d := NewDeck()
   d.AddFacts(f1, "afile")
   if d.GetCard(0).GetQuestion() != "hello there" { t.Fail() }
   if !d.GetCard(0).HasAnswer() { t.Fail() }
   if d.GetMeta(0) != nil { t.Fail() }

   testmeta := NewMeta("", "", "sm2", []string{})
   d.AddMeta(d.GetHash(0), testmeta)
   if d.GetMeta(0).IsZero() { t.Fail() }
   md := strings.Fields(d.GetMeta(0).String())
   if len(md) != 3 { t.Fail() }
   if md[1] != "0" { t.Fail() }
   if md[2] != "sm2" { t.Fail() }

   d.Forget(0)
   if d.GetMeta(0) != nil { t.Fail() }
   if !d.GetCard(0).HasAnswer() { t.Fail() }
   d.AddFacts(f1, "nofile")
   if d.Len() != 1 { t.Fail() }
   if d.GetCard(0).File != "afile" { t.Fail() }
   d.FilterOutFile("nofile")
   if d.Len() != 1 { t.Fail() }
   d.FilterOutFile("afile")
   if d.Len() != 0 { t.Fail() }
   d.Del(0)
   if d.GetCard(0) != nil { t.Fail() }
}

func TestDeckMove(t *testing.T) {
   d := NewDeck()
   d.AddFacts(f1, "afile")
   d.AddFacts(f2, "afile")
   d.AddFacts(f3, "afile")
   d.AddFacts(f4, "afile")
   d.TopToEnd()
   if d.TopCard().GetQuestion() != "hello" { panic("Bad moves") }
   if d.Len() != 4 { panic("Bad len") }
   d.TopTo(10)
   if d.TopCard().GetQuestion() != "i'm um" { panic("Bad moves") }
   d.TopTo(1)
   if d.TopCard().GetQuestion() != "alan the great" { panic("Bad moves") }
   d.TopTo(1)
   if d.Len() != 4 { panic("Bad len") }
   if d.TopCard().GetQuestion() != "i'm um" { panic("Bad moves") }
   d.TopToEnd()
   d.TopToEnd()
   if d.TopCard().GetQuestion() != "hello there" { panic("Bad moves") }
}