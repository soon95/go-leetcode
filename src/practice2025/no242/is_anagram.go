package no242

func isAnagram(s string, t string) bool {

	sWordCnt, tWordCnt := make(map[int32]int), make(map[int32]int)

	for _, ch := range s {
		sWordCnt[ch]++
	}

	for _, ch := range t {
		tWordCnt[ch]++
	}

	if len(sWordCnt) != len(tWordCnt) {
		return false
	}

	for key, value := range sWordCnt {

		if tWordCnt[key] != value {
			return false
		}

	}
	return true

}
