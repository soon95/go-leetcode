package main

import (
	"fmt"
	"strconv"
)

/*
*394. 字符串解码
https://leetcode.cn/problems/decode-string/
*/
func decodeString(s string) string {

	if len(s) == 0 {
		return ""
	}

	if s[0] >= 'a' && s[0] <= 'z' {
		return s[0:1] + decodeString(s[1:])
	}

	// 打头的是数字

	index := 0
	for s[index] != '[' {
		index++
	}

	temp := 1
	endIndex := index + 1
	for temp > 0 {
		if s[endIndex] == '[' {
			temp += 1
		} else if s[endIndex] == ']' {
			temp -= 1
		}
		endIndex++
	}

	num, _ := strconv.Atoi(s[:index])
	str := decodeString(s[index+1 : endIndex-1])

	ans := ""
	for i := 0; i < num; i++ {
		ans += str
	}

	return ans + decodeString(s[endIndex:])
}

func main() {

	//s := "3[a]2[bc]"
	//s := "2[abc]3[cd]ef"
	s := "3[a2[c]]"

	fmt.Printf("%v", decodeString(s))

}
