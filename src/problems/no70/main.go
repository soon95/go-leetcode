package main

import "fmt"

/*
*70. 爬楼梯
https://leetcode.cn/problems/climbing-stairs/
*/
func climbStairs(n int) int {

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	dp := make([]int, n)

	dp[0] = 1
	dp[1] = 2

	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n-1]
}

func main() {

	n := 3

	fmt.Print(climbStairs(n))

}
