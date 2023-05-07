package main

import "fmt"

/*
*
48. 旋转图像
https://leetcode.cn/problems/rotate-image/
*/
func rotate(matrix [][]int) {

	length := len(matrix)

	for i := 0; i < length/2; i++ {
		for j := 0; j < (length+1)/2; j++ {

			temp := matrix[i][j]

			matrix[i][j] = matrix[length-1-j][i]
			matrix[length-1-j][i] = matrix[length-1-i][length-1-j]
			matrix[length-1-i][length-1-j] = matrix[j][length-1-i]
			matrix[j][length-1-i] = temp

		}
	}

}

func main() {

	//matrix := [][]int{
	//	{1, 2, 3},
	//	{4, 5, 6},
	//	{7, 8, 9},
	//}

	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	rotate(matrix)

	fmt.Printf("%v", matrix)

}
