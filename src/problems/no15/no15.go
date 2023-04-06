package main

import (
	"fmt"
	"sort"
)

/*
15. 三数之和
https://leetcode.cn/problems/3sum/
*/
func threeSum(nums []int) [][]int {

	result := [][]int{}

	length := len(nums)

	sort.Ints(nums)

	start := 0

	var middle int
	var end int

	for start < length-2 {

		middle = start + 1
		end = length - 1

		if nums[start] > 0 {
			break
		}
		if nums[end] < 0 {
			break
		}

		for middle < end {

			sum := nums[start] + nums[middle] + nums[end]

			if sum == 0 {
				// 找到了

				if len(result) == 0 ||
					result[len(result)-1][0] != nums[start] ||
					result[len(result)-1][1] != nums[middle] ||
					result[len(result)-1][2] != nums[end] {

					temp := make([]int, 3)
					temp[0] = nums[start]
					temp[1] = nums[middle]
					temp[2] = nums[end]

					if !exist(temp, result) {
						result = append(result, temp)
					}
				}

				end--
				middle++

			} else if sum > 0 {
				end--
			} else {
				middle++
			}
		}

		start++

	}

	return result
}

func exist(temp []int, result [][]int) bool {

	for _, r := range result {

		if r[0] == temp[0] && r[1] == temp[1] && r[2] == temp[2] {
			return true
		}
	}
	return false
}

func main() {

	nums := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}

	result := threeSum(nums)

	fmt.Printf("%v\n", result)

}
