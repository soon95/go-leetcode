package main

import "fmt"

/*
75. 颜色分类
*https://leetcode.cn/problems/sort-colors/
*/
func sortColors(nums []int) {

	zeroCnt := 0
	oneCnt := 0

	for _, num := range nums {

		if num == 0 {
			zeroCnt++
		} else if num == 1 {
			oneCnt++
		}
	}

	for i := 0; i < len(nums); i++ {

		if zeroCnt != 0 {
			nums[i] = 0
			zeroCnt--
		} else if oneCnt != 0 {
			nums[i] = 1
			oneCnt--
		} else {
			nums[i] = 2
		}

	}

}

func main() {

	nums := []int{2, 1}

	sortColors(nums)

	fmt.Printf("%v", nums)

}
