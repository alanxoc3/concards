package internal

import (
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

func (h Hash) String() string {
	return fmt.Sprintf("%x", [16]byte(h))
}

type RKey struct {
	Hash  Hash
	Total int
}
