package no153

import "math"

func findMin(nums []int) int {

	minNum := math.MaxInt

	start, end := 0, len(nums)-1

	for start <= end {

		minNum = min(minNum, nums[start])

		if nums[start] < nums[end] {
			break
		}

		mid := (start + end) / 2
		minNum = min(minNum, nums[mid])

		if nums[start] <= nums[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return minNum
}
