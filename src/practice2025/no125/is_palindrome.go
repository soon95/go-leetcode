package no125

func isPalindrome(s string) bool {

	chars := make([]int32, 0)
	for _, ch := range s {
		if (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') {
			chars = append(chars, ch)
		} else if ch >= 'A' && ch <= 'Z' {
			chars = append(chars, ch+('a'-'A'))
		}
	}

	for i := 0; i < len(chars)/2; i++ {

		if chars[i] != chars[len(chars)-1-i] {
			return false
		}
	}

	return true
}
