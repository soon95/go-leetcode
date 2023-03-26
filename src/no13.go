package main

import "fmt"

/*
13. 罗马数字转整数
https://leetcode.cn/problems/roman-to-integer/
*/
func romanToInt(s string) int {

	dic := map[string]int{}
	dic["I"] = 1
	dic["V"] = 5
	dic["X"] = 10
	dic["L"] = 50
	dic["C"] = 100
	dic["D"] = 500
	dic["M"] = 1000

	index := 0

	num := 0

	for index < len(s) {
		cur := s[index : index+1]

		if index+1 < len(s) && dic[cur] < dic[s[index+1:index+2]] {

			num += dic[s[index+1:index+2]] - dic[cur]

			index = index + 2
		} else {
			num += dic[cur]
			index++
		}
	}

	return num
}

func main() {

	s := "DCXXI"

	fmt.Println(romanToInt(s))

}
