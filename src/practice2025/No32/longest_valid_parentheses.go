package No32

func longestValidParentheses(s string) int {

	leftCnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			leftCnt++
		}
	}
	if leftCnt == 0 || leftCnt == len(s) {
		return 0
	}

	ans := 0
	for i := 0; i < len(s); i++ {
		dp := make([]int, len(s))

		for j := i; j < len(s); j++ {

			if s[j] == '(' {
				dp[j] = dp[j-1] + 1
			} else {
				dp[j] = dp[j-1] - 1
			}

			if dp[j] == 0 {
				ans = max(ans, j-i+1)
			} else if dp[j] < 0 {
				break
			} else if dp[j] > (len(s)-i)/2 {
				break
			}
		}

	}

	return ans

}
