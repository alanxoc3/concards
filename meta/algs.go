package meta

type AlgFunc func(Outcome) float64
type AnswerClassification uint8

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
