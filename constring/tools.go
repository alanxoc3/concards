package constring

import "strings"

func Trim(str string) string {
	return strings.Trim(str, " \n\t")
}

func Split(str string) []string {
	return strings.Split(str, " ")
}

func StringToList(str string) *[]string {
	oldLst := Split(Trim(str))
	var newLst []string

	for i, x := range oldLst {
		oldLst[i] = Trim(x)
		if len(oldLst[i]) > 0 {
			newLst = append(newLst, oldLst[i])
		}
	}

	return &newLst
}
