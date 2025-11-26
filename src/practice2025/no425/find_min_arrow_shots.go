package no425

import "sort"

func findMinArrowShots(points [][]int) int {

	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		} else {
			return points[i][0] < points[j][0]
		}
	})

	intersection := points[0]

	ans := 1
	for i := 1; i < len(points); i++ {

		if points[i][0] > intersection[1] || points[i][1] < intersection[0] {
			ans++
			intersection = points[i]
		} else {
			intersection[0] = max(intersection[0], points[i][0])
			intersection[1] = min(intersection[1], points[i][1])
		}
	}

	return ans
}
