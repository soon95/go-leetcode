package no125

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	//s := "A man, a plan, a canal: Panama"
	s := "0P"

	res := isPalindrome(s)

	fmt.Printf("%v", res)
}
