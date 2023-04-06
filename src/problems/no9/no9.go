package main

import (
	"fmt"
	"strconv"
)

/*
*
9. 回文数
https://leetcode.cn/problems/palindrome-number/
*/
func isPalindrome(x int) bool {

	str := strconv.Itoa(x)

	length := len(str)

	for i := 0; i < length/2; i++ {

		if str[i] != str[length-1-i] {
			return false
		}
	}

	return true
}

func main() {

	x := 1

	fmt.Println(isPalindrome(x))

}
