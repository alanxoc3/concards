package constring

import "strings"
import "errors"
import "time"
import "fmt"

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

// Formats a list of strings as: x|y|z|a...
func FormatList(clist []string) string {
	outStr := ""

	for i, x := range clist {
		if i == 0 {
			outStr += x
		} else {
			outStr += "|" + x
		}
	}

	return outStr
}
