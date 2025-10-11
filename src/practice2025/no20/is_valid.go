package no20

func isValid(s string) bool {

	stack := make([]byte, 0)

	index := 0

	for index < len(s) {

		tmp := s[index]

		if tmp == '(' || tmp == '[' || tmp == '{' {
			stack = append(stack, tmp)
		} else {

			if len(stack) == 0 || (tmp == ')' && stack[len(stack)-1] != '(') || (tmp == ']' && stack[len(stack)-1] != '[') || (tmp == '}' && stack[len(stack)-1] != '{') {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}

		index++
	}

	return len(stack) == 0
}
