package main

import (
	"fmt"
)

/*
*
34. 在排序数组中查找元素的第一个和最后一个位置
https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
*/
func searchRange(nums []int, target int) []int {

	targetStart := -1
	targetEnd := -1

	if len(nums) == 0 || nums[0] > target || nums[len(nums)-1] < target {
		return []int{targetStart, targetEnd}
	}

	preStart := 0
	preEnd := len(nums) - 1

	for preStart <= preEnd {

		middle := (preStart + preEnd) / 2

		if nums[middle] < target {

			if nums[middle+1] == target {
				targetStart = middle + 1
				break
			} else if nums[middle+1] > target {
				break
			}

			preStart = middle

		} else if nums[middle] == target && middle == 0 {
			targetStart = middle
			break
		} else {
			preEnd = middle
		}

	}

	postStart := 0
	postEnd := len(nums) - 1

	for postStart <= postEnd {

		middle := (postStart + postEnd + 1) / 2

		if nums[middle] > target {

			if nums[middle-1] == target {
				targetEnd = middle - 1
				break
			} else if nums[middle-1] < target {
				break
			}

			postEnd = middle

		} else if nums[middle] == target && middle == len(nums)-1 {
			targetEnd = middle
			break
		} else {
			postStart = middle
		}

	}

	return []int{targetStart, targetEnd}
}

func searchRangeV2(nums []int, target int) []int {

	start := doSearch(nums, target)

	end := doSearch(nums, target+1)

	startAns := -1
	endAns := -1

	if start < len(nums) && nums[start] == target {
		startAns = start
	}

	if end > 0 && nums[end-1] == target {
		endAns = end - 1
	}

	return []int{startAns, endAns}
}

/*
*
寻找 >=target 的第一个数
这个是核心
*/
func doSearch(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {

		middle := (start + end) / 2

		if nums[middle] < target {

			start = middle + 1

		} else {

			end = middle - 1
		}

	}

	return start
}

func main() {

	nums := []int{}
	target := 5

	fmt.Printf("%v\n", searchRangeV2(nums, target))

}
