package main

import "fmt"

/*
*461. 汉明距离
https://leetcode.cn/problems/hamming-distance/
*/
func hammingDistance(x int, y int) int {

	ans := 0

	temp := x ^ y

	for temp > 0 {

		if temp&1 == 1 {
			ans++
		}
		temp >>= 1

	}

	return ans
}

func main() {

	//x, y := 1, 4
	x, y := 3, 1

	fmt.Printf("%v", hammingDistance(x, y))

}
