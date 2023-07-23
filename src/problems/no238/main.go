package main

import "fmt"

/*
*238. 除自身以外数组的乘积
https://leetcode.cn/problems/product-of-array-except-self/
*/
func productExceptSelf(nums []int) []int {

	length := len(nums)

	headTemp := make([]int, length)
	headTemp[0] = 1
	tailTemp := make([]int, length)
	tailTemp[length-1] = 1

	for i := 1; i < length; i++ {

		headTemp[i] = headTemp[i-1] * nums[i-1]

		tailTemp[length-1-i] = tailTemp[length-i] * nums[length-i]
	}

	ans := make([]int, length)

	for i := 0; i < length; i++ {

		ans[i] = headTemp[i] * tailTemp[i]
	}

	return ans
}

func main() {

	//nums := []int{100, 2, 3, 4}
	nums := []int{-1, 1, 0, -3, 3}

	fmt.Printf("%v", productExceptSelf(nums))

}
