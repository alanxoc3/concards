package main

import (
	"fmt"
	"testing"

	"github.com/alanxoc3/concards-go/constring"
)

func TestCon1(t *testing.T) {
	arr := constring.StringToList(" o a o asn taeoh saoenut ba  aoe eoth oeu   a eo a   ao")
	fmt.Println(arr)
	fmt.Println(len(*arr))

	tryToScan := "   20-2-2017@21:57 "
	dat, err := constring.StrToDate(tryToScan)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(dat)
	fmt.Printf("Date as string: \"%s\"\n", constring.DateToString(dat))
	fmt.Println()
	fmt.Println("TESTING THE STRING LIST THING: ")
	lst := []string{}
	lst = append(lst, "umm")
	lst = append(lst, "okay")
	lst = append(lst, "I like")
	lst = append(lst, "this test")
	fmt.Println(constring.FormatList(lst))
	fmt.Println()
}
