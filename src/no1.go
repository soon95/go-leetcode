package main

import (
	"fmt"
)

/*
*
两数之和
https://leetcode.cn/problems/two-sum/
*/
func twoSum(nums []int, target int) []int {

	set := make(map[int]int)

	for i, num := range nums {
		set[num] = i
	}

	for index, num := range nums {

		part := target - num

		nextIndex, exist := set[part]

		if exist && index != nextIndex {
			return []int{index, nextIndex}
		}

	}

	return nil
}

func main() {

	nums := []int{3, 2, 4}
	target := 6

	sum := twoSum(nums, target)

	fmt.Print("结果为:%v", sum)
}
