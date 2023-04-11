package main

import "fmt"

var ans []string

/*
22. 括号生成
*https://leetcode.cn/problems/generate-parentheses/
*/
func generateParenthesis(n int) []string {

	ans = []string{}

	cur := "("
	doGenerateParenthesis(cur, 1, 0, n)

	return ans
}

func doGenerateParenthesis(pre string, left int, right int, n int) {

	if right == n {
		ans = append(ans, pre)
		return
	}

	if left == right {
		doGenerateParenthesis(pre+"(", left+1, right, n)

	} else if left == n {
		// 如果左括号满了，只能加右括号了
		doGenerateParenthesis(pre+")", left, right+1, n)
	} else {
		// 如果左括号没满，既可以加左括号，又可以加右括号
		doGenerateParenthesis(pre+"(", left+1, right, n)
		doGenerateParenthesis(pre+")", left, right+1, n)

	}

}

func main() {

	n := 1

	fmt.Printf("%v\n", generateParenthesis(n))

}
