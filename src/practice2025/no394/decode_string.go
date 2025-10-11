package no394

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {

	if len(s) == 0 {
		return ""
	}

	if s[0] >= 'a' && s[0] <= 'z' {
		return string(s[0]) + decodeString(s[1:])
	}

	index, start, end, leftCnt := 0, 0, 0, 0
	for index < len(s) {

		if s[index] == '[' {

			if leftCnt == 0 {
				start = index
			}

			leftCnt++
		} else if s[index] == ']' {

			if leftCnt == 1 {
				end = index
				break
			}

			leftCnt--
		}

		index++
	}

	times, _ := strconv.Atoi(s[:start])

	return strings.Repeat(decodeString(s[start+1:end]), times) + decodeString(s[end+1:])
}
