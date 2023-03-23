package main

import (
	"fmt"
	"strings"
)

/*
12. 整数转罗马数字
https://leetcode.cn/problems/integer-to-roman/
*/
func intToRoman(num int) string {

	roman := ""

	scaleList := []int{1, 5, 10, 50, 100, 500, 1000}
	romanList := []string{"I", "V", "X", "L", "C", "D", "M"}

	index := 6

	for index >= 0 {
		times := num / scaleList[index]

		if times < 4 {
			roman = roman + strings.Repeat(romanList[index], times)
		} else if times == 4 {
			roman = roman + romanList[index] + romanList[index+1]
		} else if times < 9 {
			roman = roman + romanList[index+1] + strings.Repeat(romanList[index], times-5)
		} else if times == 9 {
			roman = roman + romanList[index] + romanList[index+2]
		}

		num = num % scaleList[index]

		index = index - 2
	}

	return roman
}

func main() {

	num := 1994

	fmt.Println(intToRoman(num))

}
