package meta

type algFunc func(Outcome) float64

// TODO: Remove this with GH-45.
var algs = map[string]algFunc{
	"sm2": sm2Exec,
}
