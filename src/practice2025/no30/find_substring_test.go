package no30

import (
	"fmt"
	"testing"
)

func Test_findSubstring(t *testing.T) {
	//s := "barfoothefoobar"
	//words := []string{"foo", "bar"}
	//s := "wordgoodgoodgoodbestword"
	//words := []string{"word", "good", "best", "word"}
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "good"}
	//s := "ababababab"
	//words := []string{"ababa", "babab"}

	res := findSubstring(s, words)

	fmt.Printf("%v", res)

}
