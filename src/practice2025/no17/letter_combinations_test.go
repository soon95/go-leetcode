package no17

import (
	"fmt"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	digits := "23"

	res := letterCombinations(digits)

	fmt.Printf("%v", res)

}
