package main

import "fmt"

var ans [][]int

/*
*
46. 全排列
https://leetcode.cn/problems/permutations/
*/
func permute(nums []int) [][]int {

	ans = [][]int{}

	doPermute([]int{}, nums)

	return ans
}

func doPermute(target []int, nums []int) {

	if len(nums) == 0 {
		ans = append(ans, target)
		return
	}

	for i := 0; i < len(nums); i++ {

		newTarget := append(target, nums[i])
		newNums := append([]int{}, nums[:i]...)
		newNums = append(newNums, nums[i+1:]...)

		doPermute(newTarget, newNums)
	}

}

func main() {

	nums := []int{1, 2, 3}

	fmt.Printf("%v", permute(nums))

}
