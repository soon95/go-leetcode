package main

import "fmt"

/*
*
no.6 N 字形变换
https://leetcode.cn/problems/zigzag-conversion/
*/
func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	length := len(s)

	dic := make([][]string, length)

	for i := 0; i < length; i++ {
		dic[i] = make([]string, numRows)
	}

	x := 0
	y := 0
	down := true

	for i := range s {

		dic[x][y] = string(s[i])

		if down {
			if y == numRows-1 {
				x++
				y--
				down = false
			} else {
				y++
			}
		} else {
			if y == 0 {
				y++
				down = true
			} else {
				y--
				x++
			}
		}
	}

	var result string

	for j := 0; j < numRows; j++ {

		for i := 0; i < length; i++ {

			s := dic[i][j]

			if s != "" {
				result += s
			}

		}

	}

	return result
}

func main() {

	s := "AB"
	numRows := 1

	//fmt.Print(s[100])

	fmt.Print(convert(s, numRows))

	// PAHNAPLSIIGYIR
	// PINALSIGYAHRPI
}
