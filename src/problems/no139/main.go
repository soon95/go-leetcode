package main

import "fmt"

/*
*139. 单词拆分
https://leetcode.cn/problems/word-break/
用动态规划也能解决这个问题 dp[j]=dp[i]&& s[i:j] in wordDict
*/
func wordBreak(s string, wordDict []string) bool {

	dp := map[string]bool{}

	return dfs(s, wordDict, dp)
}

func dfs(s string, wordDict []string, dp map[string]bool) bool {

	length := len(s)

	if length == 0 {
		return true
	}

	ans, exist := dp[s]
	if exist {
		return ans
	}

	ans = false

	for _, word := range wordDict {

		wordLength := len(word)

		if length >= wordLength && s[0:wordLength] == word && dfs(s[wordLength:], wordDict, dp) {

			ans = true

			break
		}

	}

	dp[s] = ans

	return ans
}

func main() {

	//s := "leetcode"
	//wordDict := []string{"leet", "code"}

	//s := "applepenapple"
	//wordDict := []string{"apple", "pen"}

	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}

	fmt.Printf("%v", wordBreak(s, wordDict))

}
