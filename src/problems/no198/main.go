package main

import (
	"fmt"
)

/*
*198. 打家劫舍
https://leetcode.cn/problems/house-robber/
*/
func rob(nums []int) int {

	length := len(nums)

	dp := make([]int, length+1)

	dp[0] = 0
	dp[1] = nums[0]

	for i := 1; i < length; i++ {

		temp := dp[i-1] + nums[i]
		if temp < dp[i] {
			temp = dp[i]
		}

		dp[i+1] = temp

	}

	return dp[length]
}

func main() {

	//nums := []int{2, 7, 9, 3, 1}
	nums := []int{1, 2}

	fmt.Printf("%v", rob(nums))
}
