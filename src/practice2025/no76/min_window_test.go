package no76

import (
	"fmt"
	"testing"
)

func Test_minWindow(test *testing.T) {
	//s := "ADOBECODEBANC"
	//t := "ABC"

	s := "a"
	t := "aa"

	res := minWindow(s, t)

	fmt.Printf("%v", res)

}
