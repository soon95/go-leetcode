package no15

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {

	// 首先排序

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	resMap := make(map[string][]int)

	index := 0

	for index < len(nums)-2 {

		left, right := index+1, len(nums)-1

		for left < right {
			if nums[index]+nums[left]+nums[right] == 0 {
				key := fmt.Sprintf("%d_%d_%d", nums[index], nums[left], nums[right])
				resMap[key] = []int{nums[index], nums[left], nums[right]}
				left++
				right--
			} else if nums[index]+nums[left]+nums[right] > 0 {
				right--
			} else {
				left++
			}
		}

		index++
	}

	res := make([][]int, 0)
	for _, tmp := range resMap {
		res = append(res, tmp)
	}

	return res
}
