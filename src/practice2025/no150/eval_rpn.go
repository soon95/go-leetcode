package no150

import "strconv"

func evalRPN(tokens []string) int {

	stack := make([]int, 0)

	for _, token := range tokens {

		if token == "+" || token == "-" || token == "*" || token == "/" {
			left, right := stack[len(stack)-2], stack[len(stack)-1]

			stack = stack[:len(stack)-2]

			tmp := 0
			if token == "+" {

				tmp = left + right

			} else if token == "-" {
				tmp = left - right
			} else if token == "*" {
				tmp = left * right
			} else if token == "/" {
				tmp = left / right
			}
			stack = append(stack, tmp)
		} else {
			num, _ := strconv.Atoi(token)

			stack = append(stack, num)
		}
	}

	return stack[0]
}
