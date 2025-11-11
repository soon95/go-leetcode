package no209

import "math"

func minSubArrayLen(target int, nums []int) int {

	sum := 0

	left, right := 0, 0
	ans := math.MaxInt
	for right < len(nums) {

		sum += nums[right]

		for left <= right && sum >= target {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}

		right++
	}

	if ans == math.MaxInt {
		return 0
	} else {
		return ans
	}
}
