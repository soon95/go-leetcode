package main

import (
	"fmt"
	"math"
)

/*
*76. 最小覆盖子串
https://leetcode.cn/problems/minimum-window-substring/
*/
func minWindow(s string, t string) string {

	source := make(map[string]int)

	for i := 0; i < len(t); i++ {
		cur := t[i : i+1]
		source[cur] = source[cur] + 1
	}

	start := 0
	end := 0

	target := make(map[string]int)

	ans := math.MaxInt
	ansStart := 0
	ansEnd := 0

	for end < len(s) {

		cur := s[end : end+1]

		target[cur] = target[cur] + 1

		for match(target, source) && start <= end {

			if end-start+1 < ans {
				ans = end - start + 1
				ansStart = start
				ansEnd = end
			}

			cur = s[start : start+1]
			target[cur] = target[cur] - 1

			start++
		}

		end++
	}

	if ans == math.MaxInt {
		return ""
	} else {
		return s[ansStart : ansEnd+1]
	}

}

func match(target map[string]int, source map[string]int) bool {

	for s, i := range source {
		if target[s] < i {
			return false
		}
	}
	return true
}

func main() {

	//s, t := "ADOBECODEBANC", "ABC"
	//s, t := "BANC", "ABC"
	s, t := "a", "aa"

	fmt.Printf("%v", minWindow(s, t))

}
