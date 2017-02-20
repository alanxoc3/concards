package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alanxoc3/concards-go/gui"
)

func testWithStr(str string) int {
	array := strings.Split(str, " ")

	cfg, err := gui.ParseConfig(array)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return 1
	}

	cfg.Print()
	return 0
}

func TestGui1(t *testing.T) {
	testWithStr("home apple pear --no-main ap ap")
	testWithStr("home -hrdmg apple pear --no-main")
	testWithStr("home -v -n 4 nion -g 3ha -g 4ha and 你好mye")
}
