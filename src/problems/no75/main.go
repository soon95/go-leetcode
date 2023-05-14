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

func sortColorsV2(nums []int) {

	oneStart := 0
	twoStart := 0

	for i := 0; i < len(nums); i++ {

		if nums[i] == 0 {

			if oneStart == twoStart {
				exchange(nums, i, oneStart)
			} else {
				exchange(nums, i, oneStart)
				exchange(nums, i, twoStart)
			}

			oneStart++
			twoStart++
		} else if nums[i] == 1 {
			exchange(nums, i, twoStart)
			twoStart++
		}

	}
}

func exchange(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func main() {

	//nums := []int{2, 1}
	nums := []int{2, 0, 2, 1, 1, 0}

	sortColorsV2(nums)

	fmt.Printf("%v", nums)

}
