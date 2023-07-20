package main

import "fmt"

/*
*https://leetcode.cn/problems/maximal-square/
221. 最大正方形
竟然可以用动态规划来解决   dp[i][j]=min(dp[i-1][j-1],dp[i-1][j],dp[i][j-1])+1
*/
func maximalSquare(matrix [][]byte) int {

	m := len(matrix)
	n := len(matrix[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
		} else {
			dp[0][i] = 0
		}
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = dp[i-1][j] + 1
			} else {
				dp[i][j] = 0
			}
		}
	}

	ans := 0
	for i := 0; i < m; i++ {
		temp := findMaxSquare(dp[i])
		ans = max(ans, temp)
	}

	return ans
}

func findMaxSquare(nums []int) int {

	ans := 0

	heights := make([]int, len(nums)+2)
	for i := 0; i < len(nums); i++ {
		heights[i+1] = nums[i]
	}

	indexStack := []int{}

	for i := 0; i < len(heights); i++ {

		for len(indexStack) != 0 && heights[indexStack[len(indexStack)-1]] > heights[i] {

			height := heights[indexStack[len(indexStack)-1]]

			indexStack = indexStack[:len(indexStack)-1]

			leftIndex := indexStack[len(indexStack)-1]
			rightIndex := i

			width := rightIndex - leftIndex - 1

			temp := min(height, width)

			ans = max(ans, temp*temp)

		}

		indexStack = append(indexStack, i)

	}

	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func main() {
	//matrix := [][]byte{
	//	{'1', '0', '1', '0', '0'},
	//	{'1', '0', '1', '1', '1'},
	//	{'1', '1', '1', '1', '1'},
	//	{'1', '0', '0', '1', '0'},
	//}

	matrix := [][]byte{
		{'0', '1'},
		{'1', '0'},
	}

	fmt.Printf("%v", maximalSquare(matrix))

	//nums := []int{3, 1, 3, 2, 2}
	//fmt.Printf("%v", findMaxSquare(nums))

}
