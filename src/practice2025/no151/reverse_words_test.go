package no151

import (
	"fmt"
	"testing"
)

func Test_reverseWords(t *testing.T) {
	//s := "the sky is blue"
	s := "  hello world  "

	res := reverseWords(s)

	fmt.Printf("%v", res)

}
