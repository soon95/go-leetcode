package no118

func generate(numRows int) [][]int {

	ans := make([][]int, 0)

	for i := 1; i <= numRows; i++ {
		ans = append(ans, make([]int, i))
	}

	for _, items := range ans {
		items[0] = 1
		items[len(items)-1] = 1
	}

	for i, items := range ans {

		for j := 1; j < len(items)-1; j++ {
			items[j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}

	return ans
}
