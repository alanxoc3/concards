package file

import "testing"
import "strings"
import "github.com/alanxoc3/concards/core"

const f1 = "@> hello there | i'm a beard <@"
const c1 = "c435597dd9718c64b135087e944fd614 2020-01-01T00:00:00Z 3 sm2 2.5"

const f2 = " @> hi | hello @> yoyo man go <@"
const c2 = "19d480d8751e75494e70e65c18cc7fac 2020-01-11T00:00:00Z 2 sm2 .00001"

const c3 = `19d480d8751e75494e70e65c18cc7fac 2020-01-01T00:00:00Z 2 sm2 .00001
c435597dd9718c64b135087e944fd614 2020-01-11T00:00:00Z 3 sm2 .05
b718c81a83d82bb83f82b0a8b18bb82b 2020-01-11T00:00:00Z 27 sm2 .05
`

func TestReadMetasToDeck(t *testing.T) {
	d := core.NewDeck()
	ReadCardsToDeckHelper(strings.NewReader(f1+f2), d, "")
	ReadMetasToDeckHelper(strings.NewReader(c1), d)

	for i := 0; i < d.Len(); i++ {
		_, c, m := d.Get(i)
		switch i {
		case 0:
			if c.GetFactRaw(0) != "hello there" {
				t.Fail()
			}
			if m.NextStr() != "2020-01-01T00:00:00Z" {
				t.Fail()
			}
			if m.Streak != 3 {
				t.Fail()
			}
		case 1:
		}
	}
}

func TestReadCardsToDeck(t *testing.T) {
	d := core.NewDeck()
	ReadCardsToDeckHelper(strings.NewReader(f2), d, "nihao")

	for i := 0; i < d.Len(); i++ {
		_, c, _ := d.Get(i)
		switch i {
		case 0:
			if c.GetFactRaw(0) != "hi" {
				t.Fail()
			}
			if c.GetFile() != "nihao" {
				t.Fail()
			}
		case 1:
			if c.GetFactRaw(0) != "yoyo man go" {
				t.Fail()
			}
		}
	}
}

func TestWriteMetasToString(t *testing.T) {
	d := core.NewDeck()
	ReadMetasToDeckHelper(strings.NewReader(c3), d)
	str := WriteMetasToString(d)
	a := strings.Split(str, "\n")
	b := strings.Split(c3, "\n")

	if a[0] != b[0] {
		t.Fail()
	}
	if a[1] != b[2] {
		t.Fail()
	}
	if a[2] != b[1] {
		t.Fail()
	}
}

/*
// This is a manual test.
func TestFile(t *testing.T) {
   d := core.NewDeck()
   ReadCardsToDeck(d, "../")

   for i := 0; i < d.Len(); i++ {
      _, c, _ := d.Get(i)
      println(c.String())
	}
}
*/
