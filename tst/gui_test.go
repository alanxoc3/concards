package main

import (
	"fmt"
	"testing"
	"strings"

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
	testWithStr("home apple pear --no-main")
}

func TestGui2(t *testing.T) {
}
