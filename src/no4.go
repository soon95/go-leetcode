package main

import (
	"fmt"
)

/*
no.4 寻找两个正序数组的中位数
https://leetcode.cn/problems/median-of-two-sorted-arrays/
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	totalLength := len(nums1) + len(nums2)

	index1 := 0
	index2 := 0
	count := 0

	var cur int
	var pre int

	for count <= totalLength/2 {

		pre = cur

		if index2 >= len(nums2) {
			// 如果第二个数据到头了 则直接后移第一个数组
			cur = nums1[index1]
			index1++
		} else if index1 >= len(nums1) {
			// 如果第一个数据到头了 则直接后移第二个数组
			cur = nums2[index2]
			index2++
		} else {
			// 如果两个数组都没到头 那么比较一下

			num1 := nums1[index1]
			num2 := nums2[index2]

			if num1 < num2 {
				cur = num1
				index1++
			} else {
				cur = num2
				index2++
			}
		}

		count++
	}

	if totalLength%2 == 0 {
		return float64(cur+pre) / 2
	} else {
		return float64(cur)
	}
}

func main() {

	nums1 := []int{}
	nums2 := []int{2, 4}

	median := findMedianSortedArrays(nums1, nums2)

	fmt.Print(median)
}
