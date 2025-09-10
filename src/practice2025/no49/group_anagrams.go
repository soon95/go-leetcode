package no49

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {

	alphaMap := make(map[string][]string)

	for _, str := range strs {
		temp := strSort(str)
		alphaMap[temp] = append(alphaMap[temp], str)
	}

	res := make([][]string, 0)

	for _, strings := range alphaMap {
		res = append(res, strings)
	}
	return res
}

func strSort(str string) string {

	list := make([]int32, 0)
	for _, i2 := range str {
		list = append(list, i2)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i] <= list[j]
	})

	return string(list)
}
