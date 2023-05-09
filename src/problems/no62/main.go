package main

import "fmt"

/*
*62. 不同路径
https://leetcode.cn/problems/unique-paths/
*/
func uniquePaths(m int, n int) int {

	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[m-1][i] = 1
	}

	for i := 0; i < m; i++ {
		dp[i][n-1] = 1
	}

	for i := m - 2; i >= 0; i-- {

		for j := n - 2; j >= 0; j-- {

			dp[i][j] = dp[i+1][j] + dp[i][j+1]
		}
	}

	return dp[0][0]
}

func main() {

	m, n := 3, 3

	fmt.Print(uniquePaths(m, n))

}
