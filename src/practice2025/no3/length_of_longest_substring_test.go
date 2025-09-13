package no3

import (
	"fmt"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	//s := "abcabcbb"
	//s := "bbbbb"
	s := "pwwkew"

	res := lengthOfLongestSubstring(s)

	fmt.Printf("%v", res)
}
