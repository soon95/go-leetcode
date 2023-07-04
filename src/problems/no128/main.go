package main

import (
	"fmt"
)

/*
*128. 最长连续序列
https://leetcode.cn/problems/longest-consecutive-sequence/
*/
func longestConsecutive(nums []int) int {

	dic := map[int]bool{}

	for _, num := range nums {
		dic[num] = false
	}

	length := 0

	for num, visited := range dic {

		if visited {
			continue
		}

		temp := 1
		left := num - 1
		for true {

			_, exist := dic[left]

			if exist {
				temp++
			} else {
				break
			}

			dic[left] = true

			left = left - 1
		}

		right := num + 1
		for true {

			_, exist := dic[right]

			if exist {
				temp++
			} else {
				break
			}

			dic[right] = true

			right = right + 1
		}

		if length < temp {
			length = temp
		}

	}

	return length

}

func main() {

	//nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	nums := []int{100, 4, 200, 1, 3, 2}

	fmt.Printf("%v", longestConsecutive(nums))

}
