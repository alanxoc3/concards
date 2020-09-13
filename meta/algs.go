package meta

import (
	"fmt"
	"time"
)

// Constants
const MaxNextDuration float64 = float64(time.Hour * 24 * 365 * 100) // 100 years.
const MaxYesNoStreak = 1 << 29                                      // About 500 million

// Types
type AnswerClassification uint8
type AlgFunc func(Outcome) float64
type Hash [16]byte

type RKey struct {
	Hash  Hash
	Total int
}

const (
	YesWasYes AnswerClassification = 1 << iota
	YesWasNo
	NoWasYes
	NoWasNo
)

// TODO: Remove this with GH-45.
var algs = map[string]AlgFunc{
	"sm2": sm2Exec,
}

func (h Hash) String() string {
	return fmt.Sprintf("%x", [16]byte(h))
}
