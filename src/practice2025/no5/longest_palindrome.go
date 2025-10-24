package no5

func longestPalindrome(s string) string {

	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	ans := ""
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
		if dp[i][i] && i-i+1 > len(ans) {
			ans = s[i : i+1]
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] && (i+1 > j-1 || dp[i+1][j-1]) {
				dp[i][j] = true
			}
			if dp[i][j] && j-i+1 > len(ans) {
				ans = s[i : j+1]
			}
		}

	}
	return ans
}
