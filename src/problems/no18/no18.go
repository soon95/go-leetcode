package main

import (
	"fmt"
	"sort"
)

/*
*
18. 四数之和
https://leetcode.cn/problems/4sum/
*/
func fourSum(nums []int, target int) [][]int {

	sort.Ints(nums)

	ans := [][]int{}

	for first := 0; first < len(nums); first++ {

		// 用于去重
		if first > 0 && nums[first-1] == nums[first] {
			continue
		}

		for second := first + 1; second < len(nums); second++ {

			// 用于去重
			if second > first+1 && nums[second-1] == nums[second] {
				continue
			}

			third := second + 1
			forth := len(nums) - 1

			for third < forth {

				temp := nums[first] + nums[second] + nums[third] + nums[forth]
				if temp == target {
					ans = append(ans, []int{nums[first], nums[second], nums[third], nums[forth]})

					third++
					// 用于去重
					for third < len(nums) && nums[third-1] == nums[third] {
						third++
					}

					forth--
					// 用于去重
					for forth >= 0 && nums[forth] == nums[forth+1] {
						forth--
					}

				} else if temp < target {

					third++

				} else {
					forth--
				}

			}

		}

	}

	return ans
}

func main() {

	nums := []int{2, 2, 2, 2, 2}
	target := 8

	fmt.Print("v%\n", fourSum(nums, target))

}
