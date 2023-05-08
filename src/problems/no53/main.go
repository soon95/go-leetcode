package main

import (
	"fmt"
	"math"
)

/*
*
53. 最大子数组和
https://leetcode.cn/problems/maximum-subarray/
内核是动态规划，但是可以简化
*/
func maxSubArray(nums []int) int {

	ans := math.MinInt

	temp := 0

	for _, num := range nums {

		temp += num

		if ans < temp {
			ans = temp
		}

		if temp < 0 {
			temp = 0
		}
	}

	return ans
}

func main() {

	nums := []int{5, 4, -1, 7, 8}

	fmt.Print(maxSubArray(nums))

}
