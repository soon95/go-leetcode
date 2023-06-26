package main

import "fmt"

/*
*
79. 单词搜索
https://leetcode.cn/problems/word-search/
*/
func exist(board [][]byte, word string) bool {

	length := len(board)
	width := len(board[0])

	visited := make([][]bool, length)
	for i := 0; i < length; i++ {

		visited[i] = make([]bool, width)

	}

	for i := 0; i < length; i++ {

		for j := 0; j < width; j++ {

			if dfs(i, j, visited, board, word) {
				return true
			}
		}
	}

	return false
}

func dfs(i, j int, visited [][]bool, board [][]byte, word string) bool {

	// 如果单词为空，返回成功
	if len(word) == 0 {
		return true
	}

	length := len(board)
	width := len(board[0])

	// 如果超出边界了，返回失败
	if i < 0 || i >= length || j < 0 || j >= width {
		return false
	}

	// 如果字符访问过，返回失败
	if visited[i][j] {
		return false
	}

	// 如果字母不匹配，返回失败
	if board[i][j] != word[0] {
		return false
	}

	curVisited := make([][]bool, length)
	for k := 0; k < length; k++ {
		curVisited[k] = make([]bool, width)
		copy(curVisited[k], visited[k])
	}

	curVisited[i][j] = true

	nextWord := word[1:]

	if len(nextWord) == 0 {
		return true
	} else {

		return dfs(i+1, j, curVisited, board, nextWord) ||
			dfs(i-1, j, curVisited, board, nextWord) ||
			dfs(i, j+1, curVisited, board, nextWord) ||
			dfs(i, j-1, curVisited, board, nextWord)
	}

}

func main() {

	//board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	//word := "ABCB"
	board := [][]byte{{'C', 'A', 'A'}, {'A', 'A', 'A'}, {'B', 'C', 'D'}}
	word := "AAB"

	fmt.Printf("%v", exist(board, word))

}
