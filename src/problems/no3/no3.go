package main

import (
	"fmt"
)

/*
* no3 无重复字符的最长子串
* https://leetcode.cn/problems/longest-substring-without-repeating-characters/
 */
func lengthOfLongestSubstring(s string) int {

	length := len(s)

	if length == 0 {
		return 0
	}

	var maxLength int

	start := 0
	end := 0

	charMap := map[byte]int{}

	for end < length {

		cur := s[end]

		_, exist := charMap[cur]

		if exist {
			// 如果有重复，就删除其实位置的字符，并将起始位置向前移动一格
			duplicate := s[start]

			delete(charMap, duplicate)

			start++

		} else {
			// 如果没有重复 将终点位置的字符加入到map中，并将终点位置向前移动一格
			charMap[cur] = 0

			if len(charMap) > maxLength {
				maxLength = len(charMap)
			}

			end++
		}

	}

	return maxLength
}

func main() {

	s := "abcdefg"

	res := lengthOfLongestSubstring(s)

	fmt.Print(res)

}
