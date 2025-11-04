package no12

import (
	"fmt"
	"testing"
)

func Test_intToRoman(t *testing.T) {
	num := 3749

	res := intToRoman(num)

	fmt.Printf("%v", res)

}
