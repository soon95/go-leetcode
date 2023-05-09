package main

import (
	"fmt"
)

/*
*64. 最小路径和
https://leetcode.cn/problems/minimum-path-sum/
*/
func minPathSum(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]

	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {

			if dp[i-1][j] < dp[i][j-1] {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}

		}
	}

	return dp[m-1][n-1]
}

func main() {

	//grid := [][]int{
	//	{1, 3, 1},
	//	{1, 5, 1},
	//	{4, 2, 1},
	//}
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Print(minPathSum(grid))

}
