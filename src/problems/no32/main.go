package main

import (
	"fmt"
)

/*
*
32. 最长有效括号

通过动态规划来解题

https://leetcode.cn/problems/longest-valid-parentheses/
*/
func longestValidParentheses(s string) int {

	length := len(s)

	dp := make([]int, length)

	for i := 1; i < length; i++ {

		if s[i] == ')' {

			if s[i-1] == '(' {
				if i == 1 {
					dp[i] = 2
				} else {
					dp[i] = dp[i-2] + 2
				}
			} else if i-1-dp[i-1] >= 0 && s[i-1-dp[i-1]] == '(' {
				if i-1-dp[i-1]-1 >= 0 {
					dp[i] = dp[i-1] + 2 + dp[i-1-dp[i-1]-1]
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
	}

	ans := 0

	for i := 0; i < length; i++ {
		if dp[i] > ans {
			ans = dp[i]
		}
	}

	return ans
}

/*
*
通过栈来解题
*/
func longestValidParenthesesV2(s string) int {

	ans := 0

	stack := []int{-1}

	for i := 0; i < len(s); i++ {

		cur := s[i]

		if cur == '(' {
			stack = append(stack, i)
		} else {

			top := stack[len(stack)-1]

			if top >= 0 && s[top] == '(' {

				stack = stack[:len(stack)-1]
				top = stack[len(stack)-1]

				l := i - top

				if l > ans {
					ans = l
				}
			} else {

				stack = append(stack, i)

			}

		}

	}

	return ans
}

func main() {

	s := "(()"

	fmt.Print(longestValidParenthesesV2(s))

}
