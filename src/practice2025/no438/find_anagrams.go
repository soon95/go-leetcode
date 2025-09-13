package no438

func findAnagrams(s string, p string) []int {

	targetMap := make(map[int32]int)

	for _, ch := range p {
		targetMap[ch] = targetMap[ch] + 1
	}

	res := make([]int, 0)

	if len(s) < len(p) {
		return res
	}

	sourceMap := make(map[byte]int)
	for i := 0; i < len(p)-1; i++ {
		sourceMap[s[i]] = sourceMap[s[i]] + 1
	}

	for i := len(p) - 1; i < len(s); i++ {

		sourceMap[s[i]] = sourceMap[s[i]] + 1

		if i-len(p) >= 0 {
			sourceMap[s[i-len(p)]] = sourceMap[s[i-len(p)]] - 1
		}

		find := true
		for ch, cnt := range targetMap {
			if cnt != sourceMap[byte(ch)] {
				find = false
			}
		}
		if find {
			res = append(res, i-len(p)+1)
		}
	}
	return res
}
