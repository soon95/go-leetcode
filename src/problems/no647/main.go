package main

import "fmt"

/*
*647. 回文子串
https://leetcode.cn/problems/palindromic-substrings/
*/
func countSubstrings(s string) int {

	length := len(s)
	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}

	for i := 0; i < length; i++ {
		dp[i][i] = true
	}

	for i := length - 2; i >= 0; i-- {

		for j := i + 1; j < length; j++ {

			if s[i] == s[j] && (i+1 > j-1 || dp[i+1][j-1]) {
				dp[i][j] = true
			}
		}
	}

	ans := 0
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if dp[i][j] {
				ans++
			}
		}
	}

	return ans
}

func main() {

	s := "aaa"

	fmt.Printf("%v", countSubstrings(s))
}
