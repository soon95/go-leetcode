package main

import "fmt"

var ans int

/*
*494. 目标和
https://leetcode.cn/problems/target-sum/
还有一种动态规划的思路，dp[i,j]：定义为0-i的数组中，和为j的数量
*/
func findTargetSumWays(nums []int, target int) int {

	ans = 0

	dfs(nums, 0, 0, target)

	return ans
}

func dfs(nums []int, curIndex, pre, target int) {

	length := len(nums)
	if length == curIndex {
		if pre == target {
			ans++
		}
		return
	}

	dfs(nums, curIndex+1, pre+nums[curIndex], target)
	dfs(nums, curIndex+1, pre-nums[curIndex], target)
}

func main() {

	nums := []int{1, 1, 1, 1, 1}
	target := 3

	fmt.Printf("%v", findTargetSumWays(nums, target))

}
