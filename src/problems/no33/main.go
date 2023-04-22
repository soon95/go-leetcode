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

func main() {

	nums := []int{4, 5, 6, 7, 0, 1, 2}

	target := 0

	fmt.Print(search(nums, target))

}
