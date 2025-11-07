package no14

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	index := 0

	for {

		match := true

		if index >= len(strs[0]) {
			match = false
		} else {
			tmp := strs[0][index]
			for i := 1; i < len(strs); i++ {
				if index >= len(strs[i]) || strs[i][index] != tmp {
					match = false
				}
			}
		}

		if match {
			index++
		} else {
			break
		}

	}

	return strs[0][:index]
}
