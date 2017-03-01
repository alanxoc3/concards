package algs

import (
	"fmt"
	"testing"
)

func TestIt(t *testing.T) {
	var item *SpaceAlg
	var err error

	item, err = GenSpaceAlg("")
	fmt.Println(item.ToString())

	item, err = GenSpaceAlg("apple")
	fmt.Println(item.ToString())

	item, err = GenSpaceAlg("apple|||")
	fmt.Println(item.ToString())

	item, err = GenSpaceAlg("apple|1-1-2017")
	fmt.Println(item.ToString())

	item, err = GenSpaceAlg("apple|1-1-2017|3-3-2018")
	fmt.Println(item.ToString())

	item, err = GenSpaceAlg("apple|1-1-2017|3-3-2018|23")
	fmt.Println(item.ToString())

	item, err = GenSpaceAlg("apple|1-1-2017|3-3-2018|23|-3")
	fmt.Println(item.ToString())

	item, err = GenSpaceAlg("apple|1-1-2017|3-3-2018|23|-3|3|and_the_rest__s__ere|")
	fmt.Println(item.ToString())

	item, err = GenSpaceAlg("apple|1-1-2017|3-3-2018|23|-3|3|an|_the|rest|_s|ere")
	fmt.Println(item.ToString())

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("--- NO ERRORS ----")
	}
}

func TestThat(t *testing.T) {
	var item *SpaceAlg
	var err error

	item, err = GenSpaceAlg("apple|1-1-2017|3-3-2018|23|-3|3|an|_the|rest|_s|ere")
	fmt.Println(item.ToString())

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("--- NO ERRORS ----")
	}
}
