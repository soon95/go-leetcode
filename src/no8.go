package main

import (
	"fmt"
	"math"
	"strings"
)

/*
8. 字符串转换整数 (atoi)
https://leetcode.cn/problems/string-to-integer-atoi/
*/
func myAtoi(s string) int {

	// 首先去除空格
	s = strings.TrimSpace(s)

	num := 0

	length := len(s)
	if length == 0 {
		return num
	}

	start := s[0]

	negative := start == '-'

	var startIndex int
	if start == '+' || start == '-' {
		startIndex = 1
	} else {
		startIndex = 0
	}

	for i := startIndex; i < length; i++ {

		c := s[i]

		if c < '0' || c > '9' {
			break
		}

		num = num*10 + int(c-'0')

		if negative && num > int(math.Pow(2, 31)) {
			num = int(math.Pow(2, 31))
			break
		} else if !negative && num > int(math.Pow(2, 31))-1 {
			num = int(math.Pow(2, 31)) - 1
			break
		}

	}

	if negative {
		return -num
	} else {
		return num
	}
}

func main() {

	s := "-1"

	fmt.Println(myAtoi(s))
}
