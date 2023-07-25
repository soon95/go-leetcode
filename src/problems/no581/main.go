package main

import (
	"fmt"
	"math"
	"sort"
)

/*
*581. 最短无序连续子数组
https://leetcode.cn/problems/shortest-unsorted-continuous-subarray/
*/
func findUnsortedSubarray(nums []int) int {

	length := len(nums)

	sortNums := make([]int, length)
	for i, num := range nums {
		sortNums[i] = num
	}

	sort.Ints(sortNums)

	startIndex, endIndex := 1, 0
	for i := 0; i < length; i++ {
		if nums[i] != sortNums[i] {
			startIndex = i
			break
		}
	}

	for i := length - 1; i >= 0; i-- {
		if nums[i] != sortNums[i] {
			endIndex = i
			break
		}
	}

	return endIndex - startIndex + 1
}

func findUnsortedSubarrayV2(nums []int) int {

	length := len(nums)

	leftMinNum, rightMaxNum := math.MaxInt, math.MinInt
	leftIndex, rightIndex := 1, 0
	for i := 0; i < length; i++ {
		if nums[i] < rightMaxNum {
			rightIndex = i
		} else {
			rightMaxNum = nums[i]
		}
	}

	for i := length - 1; i >= 0; i-- {
		if nums[i] > leftMinNum {
			leftIndex = i
		} else {
			leftMinNum = nums[i]
		}
	}

	return rightIndex - leftIndex + 1
}

func main() {

	//nums := []int{2, 6, 4, 8, 10, 9, 15}
	//nums := []int{1, 2, 3, 4}
	//nums := []int{1, 3, 2, 2, 2}
	nums := []int{2, 3, 3, 2, 4}

	fmt.Printf("%v", findUnsortedSubarrayV2(nums))
}
