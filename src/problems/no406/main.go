package main

import (
	"fmt"
	"sort"
)

/*
406. 根据身高重建队列
*https://leetcode.cn/problems/queue-reconstruction-by-height/
*/
func reconstructQueue(people [][]int) [][]int {

	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		} else {
			return people[i][0] > people[j][0]
		}
	})

	ans := [][]int{}

	for i := 0; i < len(people); i++ {

		temp := people[i]

		before := temp[1]

		right := append([][]int{}, ans[before:]...)

		ans = append(ans[:before], temp)
		ans = append(ans, right...)
	}

	return ans

}

func main() {

	people := [][]int{
		{7, 0},
		{4, 4},
		{7, 1},
		{5, 0},
		{6, 1},
		{5, 2},
	}

	fmt.Printf("%v", reconstructQueue(people))

}
