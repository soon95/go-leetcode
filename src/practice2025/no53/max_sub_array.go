package no53

import "math"

func maxSubArray(nums []int) int {

	res, tmp := math.MinInt, 0

	for _, num := range nums {

		tmp = tmp + num

		if res < tmp {
			res = tmp
		}

		if tmp < 0 {
			tmp = 0
		}

	}

	return res
}
