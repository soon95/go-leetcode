package no31

import "sort"

func nextPermutation(nums []int) {

	curIndex := len(nums) - 2

	for curIndex >= 0 {
		if nums[curIndex] < nums[curIndex+1] {
			break
		}
		curIndex--
	}

	if curIndex >= 0 {
		targetIndex := -1
		for i := len(nums) - 1; i > curIndex; i-- {

			if nums[i] > nums[curIndex] {
				targetIndex = i
				break
			}
		}

		if targetIndex >= 0 && curIndex >= 0 {
			nums[targetIndex], nums[curIndex] = nums[curIndex], nums[targetIndex]
		}
	}

	tmp := nums[curIndex+1:]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})

	return
}
