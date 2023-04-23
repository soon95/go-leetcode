package main

import (
	"fmt"
)

/*
*
33. 搜索旋转排序数组
https://leetcode.cn/problems/search-in-rotated-sorted-array/
*/
func search(nums []int, target int) int {

	start := 0
	end := len(nums) - 1

	for start <= end {

		middle := (start + end) / 2

		if nums[middle] == target {
			return middle
		} else {

			if nums[start] < nums[end] {
				// 如果是顺序的
				if nums[middle] > target {
					end = middle - 1
				} else {
					start = middle + 1
				}
			} else {

				if nums[start] > nums[middle] {
					// 断点在前面

					if nums[middle] > target {
						end = middle - 1
					} else {
						if nums[end] >= target {
							start = middle + 1
						} else {
							end = middle - 1
						}
					}
				} else {
					// 断电在后面

					if nums[middle] < target {
						start = middle + 1
					} else {

						if nums[start] <= target {
							end = middle - 1
						} else {
							start = middle + 1
						}
					}

				}
			}
		}

	}

	return -1
}

func searchV2(nums []int, target int) int {

	start := 0
	end := len(nums) - 1

	for start <= end {

		middle := (start + end) / 2

		if nums[middle] == target {
			return middle
		}

		if nums[start] <= nums[middle] {
			// 如果前面有序

			if nums[start] <= target && nums[middle] > target {
				// 如果目标在前面的有序区间内
				end = middle - 1
			} else {
				start = middle + 1
			}

		} else {
			// 如果后面有序

			if nums[middle] < target && nums[end] >= target {
				// 如果目标在后面的有序区间内
				start = middle + 1
			} else {
				end = middle - 1
			}

		}

	}

	return -1
}

func main() {

	nums := []int{4, 5, 6, 7, 0, 1, 2}

	target := 3

	fmt.Print(searchV2(nums, target))

}
