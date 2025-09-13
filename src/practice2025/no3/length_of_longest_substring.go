package no3

func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	charMap := make(map[byte]bool)

	left, right := 0, 0

	res := 0
	for left <= right && right < len(s) {

		_, exist := charMap[s[right]]

		if exist {

			delete(charMap, s[left])

			left++

		} else {
			charMap[s[right]] = true

			if res < len(charMap) {
				res = len(charMap)
			}

			right++
		}

	}

	return res
}
