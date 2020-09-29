package internal

import (
	"encoding/hex"
	"fmt"
	"time"
)

// All the keywords that concards treats special.
const CEsc = "\\"
const CSep = "|"
const CRev = ":"
const CBeg = "@>"
const CEnd = "<@"

const MaxNextDuration float64 = float64(time.Hour * 24 * 365 * 100) // 100 years.
const MaxYesNoStreak = 1 << 29                                      // About 500 million

var KeyWords = map[string]bool{
	CSep: true,
	CRev: true,
	CBeg: true,
	CEnd: true,
}

type Hash [16]byte

func NewHash(str string) (h Hash) {
	if len(str)%2 == 1 {
		str += "0"
	}
	if x, err := hex.DecodeString(str); err == nil {
		copy(h[:], x)
	}
	return h
}

func (h Hash) String() string {
	return fmt.Sprintf("%x", [16]byte(h))
}

func (h Hash) IsZero() bool {
	return [16]byte{} == [16]byte(h)
}

func AssertLogic(condition bool, message string) {
	if !condition {
		panic(fmt.Sprintf("Logic Error: %s\nPlease report this at: https://github.com/alanxoc3/concards/issues", message))
	}
}
