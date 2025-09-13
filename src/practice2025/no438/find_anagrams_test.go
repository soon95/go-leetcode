package no438

import (
	"fmt"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	//s := "cbaebabacd"
	//p := "abc"
	s := "abab"
	p := "ab"

	res := findAnagrams(s, p)

	fmt.Printf("%v", res)

}
