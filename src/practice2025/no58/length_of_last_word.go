package no58

import "strings"

func lengthOfLastWord(s string) int {

	split := strings.Split(s, " ")

	for i := len(split) - 1; i >= 0; i-- {
		if len(split[i]) != 0 {
			return len(split[i])
		}
	}

	return 0
}
