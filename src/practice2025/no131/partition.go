package no131

func partition(s string) [][]string {

	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {

			if i+1 > j-1 {

				if s[i] == s[j] {
					dp[i][j] = true
				} else {
					dp[i][j] = false
				}
			} else {
				if dp[i+1][j-1] && s[i] == s[j] {
					dp[i][j] = true
				} else {
					dp[i][j] = false
				}
			}
		}
	}

	return dfs(dp, s, 0, nil)
}

func dfs(dp [][]bool, s string, startIndex int, pre []string) [][]string {

	if startIndex >= len(s) {
		return [][]string{pre}
	}

	res := make([][]string, 0)
	for i := startIndex; i < len(s); i++ {
		if dp[startIndex][i] {
			cur := make([]string, 0)
			cur = append(cur, pre...)
			cur = append(cur, s[startIndex:i+1])
			tmp := dfs(dp, s, i+1, cur)
			res = append(res, tmp...)
		}
	}
	return res
}
