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
}
