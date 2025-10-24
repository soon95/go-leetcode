package no72

func minDistance(word1 string, word2 string) int {

	if len(word1) == 0 || len(word2) == 0 {
		return max(len(word1), len(word2))
	}

	dp := make([][]int, len(word1))
	for i := 0; i < len(word1); i++ {
		dp[i] = make([]int, len(word2))
	}
	if word1[0] != word2[0] {
		dp[0][0] = 1
	}

	for i := 1; i < len(word1); i++ {
		if word1[i] == word2[0] {
			dp[i][0] = i
		} else {
			dp[i][0] = dp[i-1][0] + 1
		}
	}
	for j := 1; j < len(word2); j++ {
		if word1[0] == word2[j] {
			dp[0][j] = j
		} else {
			dp[0][j] = dp[0][j-1] + 1
		}
	}

	for i := 1; i < len(word1); i++ {
		for j := 1; j < len(word2); j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {

				tmp := min(dp[i-1][j-1], dp[i-1][j])
				tmp = min(tmp, dp[i][j-1])
				dp[i][j] = tmp + 1
			}

		}
	}

	return dp[len(word1)-1][len(word2)-1]
}
