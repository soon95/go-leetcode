package no139

func wordBreak(s string, wordDict []string) bool {

	snapshot = make(map[string]bool)

	return doWordBreak(s, wordDict)
}

var snapshot map[string]bool

func doWordBreak(s string, wordDict []string) bool {

	if len(s) == 0 {
		return true
	}

	can, exist := snapshot[s]
	if exist {
		return can
	}

	for _, word := range wordDict {

		if len(s) >= len(word) && s[:len(word)] == word && doWordBreak(s[len(word):], wordDict) {
			snapshot[s] = true
			return true
		}
	}
	snapshot[s] = false
	return false
}
