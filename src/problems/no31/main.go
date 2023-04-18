package main

import (
	"fmt"
	"sort"
)

/*
*31. 下一个排列
https://leetcode.cn/problems/next-permutation/
*/
func nextPermutation(nums []int) {

	length := len(nums)

	maxIndex := length - 1

	curIndex := length - 1

	for curIndex >= 0 {

		if nums[curIndex] > nums[maxIndex] {
			maxIndex = curIndex
		}

		if nums[curIndex] < nums[maxIndex] {
			break
		}

		curIndex--
	}

	if curIndex >= 0 {

		// 回头找比当前元素大一点的元素
		targetIndex := length - 1
		for targetIndex > curIndex {

			if nums[targetIndex] > nums[curIndex] {
				break
			}

			targetIndex--
		}

		temp := nums[curIndex]
		nums[curIndex] = nums[targetIndex]
		nums[targetIndex] = temp
	}

	subNums := nums[curIndex+1:]
	sort.Ints(subNums)

	for i, num := range subNums {
		nums[curIndex+1+i] = num
	}

}

func main() {

	nums := []int{3, 1, 2}

	nextPermutation(nums)

	fmt.Printf("%v\n", nums)

}
