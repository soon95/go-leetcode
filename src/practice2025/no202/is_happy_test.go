package no202

import (
	"fmt"
	"testing"
)

func Test_isHappy(t *testing.T) {
	n := 19

	res := isHappy(n)

	fmt.Printf("%v", res)
}
