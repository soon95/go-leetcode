package no290

import "strings"

func wordPattern(pattern string, s string) bool {

	words := strings.Split(s, " ")

	if len(words) != len(pattern) {
		return false
	}

	fromMap, toMap := make(map[string]byte), make(map[byte]string)

	for i, word := range words {

		if fromItem, exist := fromMap[word]; exist && fromItem != pattern[i] {
			return false
		}

		if toItem, exist := toMap[pattern[i]]; exist && toItem != word {
			return false
		}

		fromMap[word] = pattern[i]
		toMap[pattern[i]] = word
	}

	return true
}
