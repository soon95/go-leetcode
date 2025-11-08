package no6

import (
	"fmt"
	"testing"
)

func Test_convert(t *testing.T) {
	//s := "PAYPALISHIRING"
	s := "AB"
	//numRows := 3
	numRows := 2

	res := convert(s, numRows)

	fmt.Printf("%v", res)

}
