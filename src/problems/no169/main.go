package main

import "fmt"

/*
*169. 多数元素
https://leetcode.cn/problems/majority-element/
*/
func majorityElement(nums []int) int {

	dic := map[int]int{}

	for _, num := range nums {
		count, exist := dic[num]

		if exist {
			dic[num] = count + 1
		} else {
			dic[num] = 1
		}
	}

	for key, value := range dic {

		if value > len(nums)/2 {

			return key
		}

	}

	return 0
}

func majorityElementV2(nums []int) int {

	candidate, count := nums[0], 1

	for i := 1; i < len(nums); i++ {

		if nums[i] == candidate {
			count++
		} else {

			count--

			if count == 0 {
				candidate = nums[i]
				count = 1
			}

		}

	}

	return candidate

}

func main() {

	//nums := []int{3, 2, 3}
	nums := []int{2, 2, 1, 1, 1, 2, 2}

	fmt.Printf("%v", majorityElementV2(nums))
}
