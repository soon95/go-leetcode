package main

import "fmt"

/*
*26. 删除有序数组中的重复项
https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
*/
func removeDuplicates(nums []int) int {

	dic := map[int]bool{}

	cur := 0

	for i := 0; i < len(nums); i++ {

		_, exist := dic[nums[i]]

		if !exist {
			nums[cur] = nums[i]
			dic[nums[i]] = true
			cur++
		}

	}

	return cur
}

func main() {

	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	k := removeDuplicates(nums)

	fmt.Printf("%v\n%v\n", nums, k)
}
