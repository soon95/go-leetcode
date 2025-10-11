package no394

import (
	"fmt"
	"testing"
)

func Test_decodeString(t *testing.T) {
	//s := "3[a]2[bc]"
	s := "3[a2[c]]"

	res := decodeString(s)

	fmt.Printf("%v", res)
}
