package no22

import (
	"fmt"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	n := 3
	res := generateParenthesis(n)

	fmt.Printf("%v", res)
}
