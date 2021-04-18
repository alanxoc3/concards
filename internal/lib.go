package internal

import (
	"encoding/hex"
	"fmt"
	"os"
	"time"
)

// All the keywords that concards treats special.
const MaxNextDuration float64 = float64(time.Hour * 24 * 365 * 100) // 100 years.
const MaxYesNoStreak = 1 << 29                                      // About 500 million

type Config struct {
	// The various true or false options
	IsVersion  bool
	IsReview   bool
	IsMemorize bool
	IsDone     bool
	IsPrint    bool
	IsFileList bool
	IsStream   bool

	Editor      string
	Number      int
	DataDir     string
	ConfigDir   string

	PredictFile string
	OutcomeFile string
	ReviewHookFile string

	Files       []string
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

func AssertLogic(condition bool, message string) {
	if !condition {
		panic(fmt.Sprintf("Logic Error: %s\nPlease report this at: https://github.com/alanxoc3/concards/issues", message))
	}
}

func AssertError(s string) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", s)
	os.Exit(1)
}
