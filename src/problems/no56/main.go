package main

import (
	"fmt"
	"sort"
)

/*
*
56. 合并区间
https://leetcode.cn/problems/merge-intervals/
*/
func merge(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		a := intervals[i]
		b := intervals[j]

		if a[0] == b[0] {
			return a[1] < b[1]
		} else {
			return a[0] < b[0]
		}
	})

	var ans [][]int

	temp := intervals[0]
	for i := 1; i < len(intervals); i++ {

		cur := intervals[i]

		if cur[0] <= temp[1] {
			// 落在里面

			if cur[1] > temp[1] {
				// 结尾更大
				temp[1] = cur[1]
			}

		} else {

			ans = append(ans, temp)

			temp = cur
		}

	}

	ans = append(ans, temp)

	return ans
}

func main() {

	intervals := [][]int{
		{1, 4},
		{0, 4},
		//{8, 10},
		//{15, 18},
	}

	fmt.Printf("%v", merge(intervals))
}
