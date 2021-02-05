package meta

import "time"

type algFunc func(Outcome) float64

// TODO: Remove this with GH-45.
var algs = map[string]algFunc{
	"sm2": sm2Exec,
	"unit-test": unitTestExec,
}

// Used for unit tests. Probably shouldn't be used for actual reviewing.
func unitTestExec(r Outcome) float64 {
    if r.Target() {
	return float64(time.Hour * 24)
    } else {
	return 0.0
    }
}
