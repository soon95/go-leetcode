package main

import (
	"fmt"
	"math"
)

/*
72. 编辑距离
*https://leetcode.cn/problems/edit-distance/
*/
func minDistance(word1 string, word2 string) int {

	m := len(word2)
	n := len(word1)

	if n == 0 {
		return m
	}

	if m == 0 {
		return n
	}

	// 动态规划数组的起始点是 "" 空字符串，这点很重要
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {

			if word2[i-1] == word1[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {

				temp := math.MaxInt
				if dp[i-1][j-1] < temp {
					temp = dp[i-1][j-1]
				}

				if dp[i-1][j] < temp {
					temp = dp[i-1][j]
				}

				if dp[i][j-1] < temp {
					temp = dp[i][j-1]
				}
				dp[i][j] = temp + 1

			}

		}
	}

	return dp[m][n]
}

func main() {

	//word1, word2 := "horse", "ros"
	//word1, word2 := "intention", "execution"
	//word1, word2 := "pneumonoultramicroscopicsilicovolcanoconiosis", "ultramicroscopically"
	word1, word2 := "pneumonou", "u"

	fmt.Print(minDistance(word1, word2))

}
