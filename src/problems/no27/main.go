package main

import "fmt"

/*
*
27. 移除元素
https://leetcode.cn/problems/remove-element/
*/
func removeElement(nums []int, val int) int {

	cur := 0

	for i := 0; i < len(nums); i++ {

		temp := nums[i]

		if temp != val {

			nums[cur] = temp
			cur++
		}
	}

	return cur
}

func main() {

	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2

	k := removeElement(nums, val)

	fmt.Printf("%v\n%v\n", nums, k)

}
