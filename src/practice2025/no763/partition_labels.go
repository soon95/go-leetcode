package no763

func partitionLabels(s string) []int {

	lastExist := make(map[int32]int)

	for i, ch := range s {
		lastExist[ch] = i
	}

	sections := make([]string, 0)

	index := 0

	for index < len(s) {

		startIndex := index

		endIndex := lastExist[int32(s[index])]

		for index <= endIndex {

			endIndex = max(endIndex, lastExist[int32(s[index])])

			index++
		}

		sections = append(sections, s[startIndex:endIndex+1])

	}

	ans := make([]int, 0)
	for _, section := range sections {
		ans = append(ans, len(section))
	}

	return ans
}
