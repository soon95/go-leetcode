package main

import (
	"fmt"
	"math"
)

/*
152. 乘积最大子数组
*https://leetcode.cn/problems/maximum-product-subarray/
*/
func maxProduct(nums []int) int {

	ans := math.MinInt

	startIndex := 0
	curIndex := 0

	negativeIndexes := []int{}
	for curIndex < len(nums) {

		if nums[curIndex] < 0 {

			negativeIndexes = append(negativeIndexes, curIndex)

		}

		if nums[curIndex] == 0 {

			if ans < 0 {
				ans = 0
			}

			temp := doMaxProduct(nums, negativeIndexes, startIndex, curIndex)
			if ans < temp {
				ans = temp
			}

			startIndex = curIndex + 1

			negativeIndexes = []int{}
		}

		curIndex++

	}

	temp := doMaxProduct(nums, negativeIndexes, startIndex, curIndex)
	if ans < temp {
		ans = temp
	}

	return ans
}

/*
*
左开右闭
*/
func doMaxProduct(nums, negativeIndexes []int, startIndex, endIndex int) int {

	ans := math.MinInt

	if startIndex == len(nums) {
		return ans
	}

	if len(negativeIndexes)%2 == 0 {
		// 如果负数是偶数

		temp := nums[startIndex]
		for i := startIndex + 1; i < endIndex; i++ {
			temp *= nums[i]
		}

		if ans < temp {
			ans = temp
		}

	} else {
		// 如果负数是奇数

		temp1 := nums[startIndex]
		for i := startIndex + 1; i < negativeIndexes[len(negativeIndexes)-1]; i++ {
			temp1 *= nums[i]
		}

		if ans < temp1 {
			ans = temp1
		}

		if negativeIndexes[0]+1 < endIndex {
			temp2 := nums[negativeIndexes[0]+1]
			for i := negativeIndexes[0] + 2; i < endIndex; i++ {
				temp2 *= nums[i]
			}
			if ans < temp2 {
				ans = temp2
			}
		}

	}

	return ans
}

/*
*
动态规划
*/
func maxProductV2(nums []int) int {

	ans := math.MinInt

	length := len(nums)

	maxDp := make([]int, length)
	minDp := make([]int, length)

	maxDp[0] = nums[0]
	minDp[0] = nums[0]

	for i := 1; i < length; i++ {

		maxTemp := math.MinInt

		if maxTemp < maxDp[i-1]*nums[i] {
			maxTemp = maxDp[i-1] * nums[i]
		}
		if maxTemp < minDp[i-1]*nums[i] {
			maxTemp = minDp[i-1] * nums[i]
		}
		if maxTemp < nums[i] {
			maxTemp = nums[i]
		}
		maxDp[i] = maxTemp

		minTemp := math.MaxInt

		if minTemp > maxDp[i-1]*nums[i] {
			minTemp = maxDp[i-1] * nums[i]
		}
		if minTemp > minDp[i-1]*nums[i] {
			minTemp = minDp[i-1] * nums[i]
		}
		if minTemp > nums[i] {
			minTemp = nums[i]
		}
		minDp[i] = minTemp

	}

	for i := 0; i < length; i++ {

		if ans < maxDp[i] {
			ans = maxDp[i]
		}

	}

	return ans
}

func main() {

	nums := []int{2, 3, -2, 4}
	//nums := []int{-2, 0, -1, 0, 2, 3, -2, 4}
	//nums := []int{-4, -3}
	//nums := []int{-2}
	//nums := []int{0}
	//nums := []int{-3, 0, 1, -2}
	//nums := []int{-4, -3, 0}

	fmt.Printf("%v", maxProductV2(nums))

}
