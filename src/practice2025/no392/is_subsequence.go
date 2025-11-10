package no392

func isSubsequence(s string, t string) bool {

	index := 0

	for _, ch := range t {

		if index >= len(s) {
			return true
		}

		if ch == int32(s[index]) {
			index++
		}
	}

	if index >= len(s) {
		return true
	} else {
		return false
	}
}
