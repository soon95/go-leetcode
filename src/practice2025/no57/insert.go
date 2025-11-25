package no57

import "sort"

func insert(intervals [][]int, newInterval []int) [][]int {

	ans := make([][]int, 0)

	ans = append(ans, intervals...)
	ans = append(ans, newInterval)

	sort.Slice(ans, func(i, j int) bool {
		if ans[i][0] == ans[j][0] {
			return ans[i][0] < ans[j][0]
		} else {
			return ans[i][0] < ans[j][0]
		}
	})

	return merge(ans)
}

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

func insertV2(intervals [][]int, newInterval []int) (ans [][]int) {
	left, right := newInterval[0], newInterval[1]
	merged := false
	for _, interval := range intervals {
		if interval[0] > right {
			// 在插入区间的右侧且无交集
			if !merged {
				ans = append(ans, []int{left, right})
				merged = true
			}
			ans = append(ans, interval)
		} else if interval[1] < left {
			// 在插入区间的左侧且无交集
			ans = append(ans, interval)
		} else {
			// 与插入区间有交集，计算它们的并集
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}
	if !merged {
		ans = append(ans, []int{left, right})
	}
	return
}
