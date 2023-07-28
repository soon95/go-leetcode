package main

import "fmt"

/*
*448. 找到所有数组中消失的数字
https://leetcode.cn/problems/find-all-numbers-disappeared-in-an-array/
*/
func findDisappearedNumbers(nums []int) []int {

	exists := make([]bool, len(nums))

	for _, num := range nums {
		exists[num-1] = true
	}

	ans := []int{}
	for i, exist := range exists {
		if !exist {
			ans = append(ans, i+1)
		}
	}

	return ans
}

func main() {

	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}

	fmt.Printf("%v", findDisappearedNumbers(nums))

}
