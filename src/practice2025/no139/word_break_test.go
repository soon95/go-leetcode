package no139

import (
	"fmt"
	"testing"
)

func Test_wordBreak(t *testing.T) {
	//s := "leetcode"
	//wordDict := []string{"leet", "code"}
	s := "catsandog"
	wordDict := []string{"cats", "dog", "sand", "and", "cat"}

	res := wordBreak(s, wordDict)

	fmt.Printf("%v", res)

}
