package no274

import "sort"

func hIndex(citations []int) int {

	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	ans := 0
	for i := 0; i < len(citations); i++ {

		if i+1 <= citations[i] {
			ans = i + 1
		} else {
			break
		}
	}

	return ans
}
