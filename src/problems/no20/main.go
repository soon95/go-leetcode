package main

import "fmt"

/*
*
20. 有效的括号
https://leetcode.cn/problems/valid-parentheses/
*/
func isValid(s string) bool {

	var stack []byte

	for _, ch := range s {

		var temp byte

		if len(stack) == 0 {
			temp = 0
		} else {
			temp = stack[len(stack)-1]
		}

		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, byte(ch))
		} else if ch == ')' {
			if temp != '(' {
				return false
			} else {
				stack = append(stack[:len(stack)-1])
			}
		} else if ch == ']' {
			if temp != '[' {
				return false
			} else {
				stack = append(stack[:len(stack)-1])
			}
		} else if ch == '}' {
			if temp != '{' {
				return false
			} else {
				stack = append(stack[:len(stack)-1])
			}
		}
	}

	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

func main() {

	s := "("

	fmt.Print(isValid(s))

}
