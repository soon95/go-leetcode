package no228

import (
	"fmt"
	"sort"
)

func summaryRanges(nums []int) []string {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	index := 0

	from, to := 0, 0

	ans := make([]string, 0)
	for index < len(nums) {

		if index+1 < len(nums) && nums[index]+1 == nums[index+1] {
			to++
		} else {
			tmp := ""
			if from == to {
				tmp = fmt.Sprintf("%d", nums[from])
			} else {
				tmp = fmt.Sprintf("%d->%d", nums[from], nums[to])
			}
			ans = append(ans, tmp)

			from, to = index+1, index+1
		}

		index++
	}

	return ans
}
