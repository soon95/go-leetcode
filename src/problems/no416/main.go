package main

import (
	"fmt"
	"sort"
)

/*
*416. 分割等和子集
https://leetcode.cn/problems/partition-equal-subset-sum/
*/
func canPartition(nums []int) bool {

	totalSum := 0

	sort.Ints(nums)

	for _, num := range nums {
		totalSum += num
	}

	dp := map[int]map[int]bool{}

	initMap := map[int]bool{}
	initMap[0] = true
	initMap[nums[0]] = true
	dp[0] = initMap

	for i := 1; i < len(nums); i++ {

		pre := dp[i-1]

		cur := map[int]bool{}

		for preSum, _ := range pre {

			temp := preSum + nums[i]

			cur[temp] = true
			cur[preSum] = true

			if temp*2 == totalSum {
				return true
			}
		}

		dp[i] = cur
	}

	return false
}

func main() {
	//nums := []int{1, 5, 11, 5}
	nums := []int{1, 2, 3, 5}
	//nums := []int{23, 13, 11, 7, 6, 5, 5}

	fmt.Printf("%v", canPartition(nums))

}
