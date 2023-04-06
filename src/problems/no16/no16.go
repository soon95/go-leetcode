package main

import (
	"fmt"
	"math"
	"sort"
)

/*
*
16. 最接近的三数之和
https://leetcode.cn/problems/3sum-closest/
*/
func threeSumClosest(nums []int, target int) int {

	length := len(nums)

	sort.Ints(nums)

	result := 0
	diff := math.MaxInt

	for start := 0; start < length-2; start++ {

		middle := start + 1
		end := length - 1

		for middle < end {

			temp := nums[start] + nums[middle] + nums[end]

			tempDiff := int(math.Abs(float64(temp - target)))

			if tempDiff < diff {
				result = temp
				diff = tempDiff
			}

			if temp == target {
				return target
			} else if temp > target {
				end--
			} else {
				middle++
			}
		}
	}

	return result
}

func main() {

	nums := []int{0, 0, 0}
	target := 1

	fmt.Println(threeSumClosest(nums, target))

}
