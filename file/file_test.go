package file

import "testing"
import "strings"

const f1 = "@> hello there @ i'm a beard <@"
const c1 = "c6cd355e32654cb4ba506b529ff32288 2020-01-01T00:00:00Z 3 sm2 2.5"

const f2 = " @> hi @ hello @> yoyo man go <@"
const c2 = "ab11ffa53d45453729f90b2aa6df9d65 2020-01-11T00:00:00Z 2 sm2 .00001"

func TestReadMetasToDeck(t *testing.T) {
   d := ReadCardsToDeckHelper(strings.NewReader(f1 + f2))
   d = ReadMetasToDeckHelper(strings.NewReader(c1), d)

   for i := 0; i < d.Len(); i++ {
      _, c, m := d.Get(i)
      switch i {
         case 0:
            if c.GetQuestion() != "hello there" { t.Fail() }
            if m.NextStr() != "2020-01-01T00:00:00Z" { t.Fail() }
            if m.Streak != 3 { t.Fail() }
         case 1:
      }
	}
}

func TestReadCardsToDeck(t *testing.T) {
   d := ReadCardsToDeckHelper(strings.NewReader(f2))

   for i := 0; i < d.Len(); i++ {
      _, c, _ := d.Get(i)
      switch i {
         case 0: if c.GetQuestion() != "hi" { t.Fail() }
         case 1: if c.GetQuestion() != "yoyo man go" { t.Fail() }
      }
	}
}

// TODO: Do this.
func TestWriteMetasToString(t *testing.T) {
   d := ReadCardsToDeckHelper(strings.NewReader(f2))

   for i := 0; i < d.Len(); i++ {
      _, c, _ := d.Get(i)
      switch i {
         case 0: if c.GetQuestion() != "hi" { t.Fail() }
         case 1: if c.GetQuestion() != "yoyo man go" { t.Fail() }
      }
	}
}
