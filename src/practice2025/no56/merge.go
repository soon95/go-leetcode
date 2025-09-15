package no56

import "sort"

func merge(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		} else {
			return intervals[i][0] < intervals[j][0]
		}
	})

	res := make([][]int, 0)

	tmp := []int{intervals[0][0], intervals[0][1]}

	for i := 1; i < len(intervals); i++ {

		left := intervals[i][0]
		right := intervals[i][1]

		if tmp[0] <= left && tmp[1] >= left && right >= tmp[1] {
			tmp[1] = right
		} else if left > tmp[1] {
			res = append(res, tmp)
			tmp = []int{left, right}
		}
	}

	res = append(res, tmp)

	return res
}
