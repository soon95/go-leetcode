package no151

import "strings"

func reverseWords(s string) string {

	split := strings.Split(s, " ")

	words := make([]string, 0)
	for i := len(split) - 1; i >= 0; i-- {
		if split[i] != "" {
			words = append(words, split[i])
		}
	}
	return strings.Join(words, " ")
}
