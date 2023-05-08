package main

import "fmt"

/*
*55. 跳跃游戏
https://leetcode.cn/problems/jump-game/
*/
func canJump(nums []int) bool {

	dp := make([]bool, len(nums))

	dp[0] = true

	for i := 1; i < len(nums); i++ {

		for j := i - 1; j >= 0; j-- {

			if dp[j] && i-j <= nums[j] {
				dp[i] = true
			}
		}
	}

	return dp[len(dp)-1]
}

func main() {

	nums := []int{3, 2, 1, 0, 4}

	fmt.Print(canJump(nums))

}
