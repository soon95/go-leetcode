package main

import (
	"fmt"
)

/*
*96. 不同的二叉搜索树
https://leetcode.cn/problems/unique-binary-search-trees/
*/
func numTrees(n int) int {

	nums := []int{}
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {

			temp := 0

			for k := i; k <= j; k++ {

				if k == i {
					temp += dp[k+1][j]
				} else if k == j {
					temp += dp[i][k-1]
				} else {
					temp += dp[i][k-1] * dp[k+1][j]
				}
			}

			dp[i][j] = temp
		}
	}

	return dp[0][n-1]
}

/*
*
首先想到递归，后来转化成动态规划
*/
func dfs(res []int) int {

	if len(res) <= 1 {
		return 1
	}

	total := 0

	for i, _ := range res {

		leftNum := dfs(res[0:i])

		rightNum := dfs(res[i+1:])

		total += leftNum * rightNum
	}

	return total
}

func main() {

	n := 3

	fmt.Printf("%v", numTrees(n))

}
