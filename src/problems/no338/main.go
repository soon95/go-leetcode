package main

import "fmt"

/*
*338. 比特位计数
https://leetcode.cn/problems/counting-bits/
*/
func countBits(n int) []int {

	ans := make([]int, n+1)

	ans[0] = 0

	pow := 2

	for num := 1; num <= n; num++ {

		for num/pow != 0 {
			pow *= 2
		}
		pow /= 2

		ans[num] = 1 + ans[num%pow]

	}

	return ans
}

func main() {

	//n := 2
	n := 5

	fmt.Printf("%v", countBits(n))

}
