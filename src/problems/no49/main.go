package main

import (
	"fmt"
	"sort"
)

/*
*
49. 字母异位词分组
https://leetcode.cn/problems/group-anagrams/
*/
func groupAnagrams(strs []string) [][]string {

	dic := make(map[string][]string)

	for _, str := range strs {

		b := []byte(str)

		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})

		temp := string(b)

		dic[temp] = append(dic[temp], str)
	}

	ans := [][]string{}

	for _, strList := range dic {

		ans = append(ans, strList)

	}

	return ans
}

func main() {

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	//strs := []string{"a"}

	fmt.Printf("%v", groupAnagrams(strs))

}
