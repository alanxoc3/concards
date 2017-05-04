package constring

import "strings"
import "sort"
import "errors"
import "time"
import "fmt"

func TrimLineBegin(line string, trimStr string) string {
	if len(trimStr) > len(line) {
		return ""
	}

	i := 0
	for ; i < len(trimStr); i++ {
		if trimStr[i] != line[i] {
			break
		}
	}

	return line[i:]
}

func DoesLineBeginWith(line string, test string) bool {
	if len(test) > len(line) {
		return false
	}

	for i := 0; i < len(test); i++ {
		if test[i] != line[i] {
			return false
		}
	}

	return true
}

func TabsToNewlines(str *string) string {
	newStr := ""
	for i := 0; i < len(*str); i++ {
		newStr += string((*str)[i])
		if string((*str)[i]) == "\n" {
			newStr += "\t"
		}
	}
	return newStr
}

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

func StrToDate(str string) (time.Time, error) {
	strs := strings.Split(str, "@")
	if len(strs) == 0 { // Nothing
		return time.Time{}, errors.New("The date was empty.")

	} else if len(strs) == 1 { // Only a date
		return StrToJustDate(str)
	} else if len(strs) == 2 { // Date and time
		dat, err1 := StrToJustDate(strs[0])
		dur, err2 := StrToJustTime(strs[1])

		if err1 != nil {
			return time.Time{}, err1
		} else if err2 != nil {
			return time.Time{}, err2
		}

		return dat.Add(dur), nil

	} else { // Too many @ symbols
		return time.Time{}, errors.New("The string had too many @ symbols")
	}
}

func StrToJustDate(str string) (time.Time, error) {
	var day, year int
	var month time.Month
	_, err := fmt.Sscanf(str, "%d-%d-%d", &day, &month, &year)
	if err != nil {
		return time.Time{}, errors.New("Invalid Date")
	}

	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC), nil
}

func StrToJustTime(str string) (time.Duration, error) {
	var hour, min time.Duration
	_, err := fmt.Sscanf(str, "%02d:%02d", &hour, &min)
	if err != nil {
		return 0, errors.New("Invalid Time")
	}
	return time.Minute*min + time.Hour*hour, nil
}

func DateToString(d time.Time) string {
	return JustDateToString(d) + "@" + JustTimeToString(d)
}

func JustTimeToString(d time.Time) string {
	return fmt.Sprintf("%02d:%02d", d.Hour(), d.Minute())
}

func JustDateToString(d time.Time) string {
	return fmt.Sprintf("%d-%d-%d", d.Day(), d.Month(), d.Year())
}

func IsEmpty(str string) bool {
	for char := range str {
		if char != ' ' && char != '\t' && char != '\n' {
			return false
		}
	}

	return true
}

// Formats a list of strings as: x, y, z, a...
func FormatList(clist []string) string {
	outStr := ""

	for _, x := range clist {
		outStr += ", " + x
	}

	return outStr
}

// A for some quantifier.
func IsInStrList(list1 []string, item string) bool {
	for _, str := range list1 {
		if item == str {
			return true
		}
	}
	return false
}

func StringListsIdentical(list1 []string, list2 []string) bool {
	sort.Strings(list1)
	sort.Strings(list2)

	if len(list1) == len(list2) {
		for i := 0; i < len(list1); i++ {
			if list1[i] != list2[i] {
				return false
			}
		}
	} else {
		return false
	}

	return true
}

func ListToString(list []string) string {
	if len(list) <= 0 {
		return ""
	}

	retStr := "##"

	for _, str := range list {
		retStr += " " + str
	}

	return retStr
}

func ListHasOtherList(list1 []string, list2 []string) bool {
	if len(list1) < len(list2) {
		return false
	}

	sort.Strings(list1)
	sort.Strings(list2)

	indList2 := 0

	for i := 0; i < len(list1); i++ {
		if list1[i] == list2[indList2] {
			indList2++
		}

		if indList2 == len(list2) {
			return true
		}
	}

	return false
}
