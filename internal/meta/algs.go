package meta

type algFunc func(Outcome) float64
type AnswerClassification uint8

const (
	YesWasYes AnswerClassification = 1 << iota
	YesWasNo
	NoWasYes
	NoWasNo
)

// TODO: Remove this with GH-45.
var algs = map[string]algFunc{
	"sm2": sm2Exec,
}