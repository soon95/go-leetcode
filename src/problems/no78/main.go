package main

import "fmt"

/*
*78. 子集
https://leetcode.cn/problems/subsets/
*/
func subsets(nums []int) [][]int {

	ans := [][]int{{}}

	for i := 0; i < len(nums); i++ {

		temp := [][]int{}

		for j := 0; j < len(ans); j++ {

			cur := []int{}

			cur = append(cur, ans[j]...)
			cur = append(cur, nums[i])

			temp = append(temp, cur)
		}

		ans = append(ans, temp...)

	}

	return ans
}

func main() {

	nums := []int{0}

	fmt.Printf("%v", subsets(nums))

}
