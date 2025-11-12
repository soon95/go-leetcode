package no30

func findSubstring(s string, words []string) []int {
	ans := make([]int, 0)

	wordLength := len(words[0])

	if len(s) < len(words)*wordLength {
		return ans
	}

	targetMap := make(map[string]int)
	for _, word := range words {
		targetMap[word]++
	}

	for i := 0; i < wordLength && i+len(words)*wordLength <= len(s); i++ {

		countMap := make(map[string]int)

		for j := 0; j < len(words); j++ {

			curWord := s[i+j*wordLength : i+(j+1)*wordLength]

			countMap[curWord]++
		}
		if mapEqual(countMap, targetMap) {
			ans = append(ans, i)
		}

		for j := len(words); i+(j+1)*wordLength <= len(s); j++ {
			preWord := s[i+(j-len(words))*wordLength : i+(j-len(words)+1)*wordLength]
			curWord := s[i+j*wordLength : i+(j+1)*wordLength]
			countMap[preWord]--
			if countMap[preWord] == 0 {
				delete(countMap, preWord)
			}
			countMap[curWord]++
			if mapEqual(countMap, targetMap) {
				ans = append(ans, i+j*wordLength-wordLength*(len(words)-1))
			}
		}
	}
	return ans
}

func mapEqual(source, target map[string]int) bool {

	if len(source) != len(target) {
		return false
	}

	for key, value := range source {
		if target[key] != value {
			return false
		}
	}

	return true
}
