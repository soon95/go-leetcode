package main

import "fmt"

/*
*136. 只出现一次的数字
https://leetcode.cn/problems/single-number/
*/
func singleNumber(nums []int) int {

	ans := 0

	for _, num := range nums {

		ans = ans ^ num

	}

	return ans
}

func main() {

	nums := []int{2, 2, 1}

	fmt.Printf("%v", singleNumber(nums))

}
