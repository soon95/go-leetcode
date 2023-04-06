package main

import "fmt"

/*
*
14. 最长公共前缀
https://leetcode.cn/problems/longest-common-prefix/
*/
func longestCommonPrefix(strs []string) string {

	commonPrefix := []byte{}

	if len(strs) == 1 {
		return strs[0]
	}

	index := 0

	for true {

		for _, str := range strs {
			if len(str) == index {
				return string(commonPrefix)
			}
		}

		temp := strs[0][index]

		for i := 1; i < len(strs); i++ {
			if temp != strs[i][index] {
				return string(commonPrefix)
			}
		}

		commonPrefix = append(commonPrefix, temp)
		index++
	}

	return string(commonPrefix)
}

func main() {

	strs := []string{"dog", "d", ""}

	fmt.Println(longestCommonPrefix(strs))

}
