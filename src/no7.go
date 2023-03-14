package main

import (
	"fmt"
	"math"
)

/*
7. 整数反转
https://leetcode.cn/problems/reverse-integer/
*/
func reverse(x int) int {

	num := 0

	items := []int{}

	base := 10

	res := x

	for res != 0 {

		items = append(items, res%base)

		res = res / base
	}

	for i := 0; i < len(items); i++ {

		num += int(math.Pow(10, float64(len(items)-i-1))) * items[i]

		if (num > int(math.Pow(2, 31))-1) || (num < -int(math.Pow(2, 31))) {
			return 0
		}

	}

	return num
}

func main() {

	x := 1563847412

	fmt.Println(reverse(x))

}
