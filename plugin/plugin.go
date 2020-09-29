package plugin

import "github.com/alanxoc3/concards/internal/meta"

type Outcome = meta.Outcome

func MockOutcome(strs ...string) *Outcome {
   return meta.NewOutcomeFromStrings(strs...)
}

const YesWasYes meta.AnswerClassification = meta.YesWasYes
const YesWasNo  meta.AnswerClassification = meta.YesWasNo
const NoWasYes  meta.AnswerClassification = meta.NoWasYes
const NoWasNo   meta.AnswerClassification = meta.NoWasNo
