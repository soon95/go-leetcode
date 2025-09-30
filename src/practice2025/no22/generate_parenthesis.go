package no22

func generateParenthesis(n int) []string {

	result = make([]string, 0)

	dfs(0, n, "")

	return result
}

var result []string

func dfs(half, res int, preStr string) {

	if half == 0 && res == 0 {
		result = append(result, preStr)
	}

	if half > 0 {
		dfs(half-1, res, preStr+")")
	}

	if res > 0 {
		dfs(half+1, res-1, preStr+"(")
	}
}
