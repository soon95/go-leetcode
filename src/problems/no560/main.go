package main

import (
	"fmt"
)

/*
560. 和为 K 的子数组
*https://leetcode.cn/problems/subarray-sum-equals-k/
想复杂了，暴力求解即可
*/
func subarraySum(nums []int, k int) int {

	ans := 0

	length := len(nums)

	for i := 0; i < length; i++ {

		temp := 0

		for j := i; j < length; j++ {

			temp += nums[j]

			if temp == k {
				ans++
			}
		}
	}

	return ans
}

func main() {

	//nums := []int{1, 1, 1, 0, 0, 0}
	//k := 2

	//nums := []int{1, 2, 3}
	//k := 3

	nums := []int{-1, -1, 1}
	k := 0

	fmt.Printf("%v", subarraySum(nums, k))

}
