package main

import "fmt"

/*
*
no.5 最长回文子串
https://leetcode.cn/problems/longest-palindromic-substring/
*/
func longestPalindrome(s string) string {

	var result = s[0:1]

	length := len(s)

	dic := make([][]bool, length)
	for i := range dic {
		dic[i] = make([]bool, length)
	}

	for i := 0; i < length; i++ {
		dic[i][i] = true
	}

	for end := 1; end < length; end++ {

		for start := 0; start < end; start++ {

			if end-start == 1 {
				if s[start] == s[end] {
					dic[start][end] = true

					if len(result) < end-start+1 {
						result = s[start : end+1]
					}

				}

			} else {
				if s[start] == s[end] && dic[start+1][end-1] {
					dic[start][end] = true

					if len(result) < end-start+1 {
						result = s[start : end+1]
					}
				}
			}

		}
	}

	return result
}

func main() {

	s := "babad"

	fmt.Print(longestPalindrome(s))

}
