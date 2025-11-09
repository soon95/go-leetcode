package no68

import (
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {

	ans := make([]string, 0)

	lineWord := make([]string, 0)

	for _, word := range words {

		tmp := make([]string, 0)
		tmp = append(tmp, lineWord...)
		tmp = append(tmp, word)

		if len(strings.Join(tmp, " ")) > maxWidth {
			// 处理当前行

			tmp = tmp[:len(tmp)-1]
			ans = append(ans, formatLine(tmp, maxWidth))

			lineWord = make([]string, 0)
		}

		lineWord = append(lineWord, word)
	}

	lastWord := strings.Join(lineWord, " ")
	lastWord += strings.Repeat(" ", maxWidth-len(lastWord))
	ans = append(ans, lastWord)

	return ans
}

func formatLine(words []string, maxWidth int) string {
	res := maxWidth - len(strings.Join(words, ""))

	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", res)
	} else {
		repeat := res / (len(words) - 1)
		other := res % (len(words) - 1)

		line := words[0]
		for i := 1; i < len(words); i++ {
			line += strings.Repeat(" ", repeat)
			if i <= other {
				line += " "
			}
			line += words[i]
		}
		return line
	}
}
