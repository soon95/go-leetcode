package no6

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	list := make([][]string, 0)

	index := 0

	for index < len(s) {

		str1 := s[index:min(index+numRows, len(s))]

		if len(str1) > 0 {

			tmp := make([]string, numRows)
			for i := 0; i < len(str1); i++ {
				tmp[i] = str1[i : i+1]
			}

			list = append(list, tmp)
		}

		if index+numRows >= len(s) {
			break
		}

		str2 := s[min(index+numRows, len(s)):min(index+2*numRows-2, len(s))]

		if len(str2) > 0 {
			tmpList := make([][]string, 0)
			for i := 0; i < len(str2); i++ {

				tmp := make([]string, numRows)

				tmp[numRows-2-i] = str2[i : i+1]

				tmpList = append(tmpList, tmp)
			}

			list = append(list, tmpList...)
		}

		index = min(index+2*numRows-2, len(s))
	}

	ans := ""

	for j := 0; j < numRows; j++ {

		for i := 0; i < len(list); i++ {
			if list[i][j] != "" {
				ans += list[i][j]
			}
		}
	}

	return ans
}
