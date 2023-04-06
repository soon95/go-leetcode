package main

import "fmt"

/*
17. 电话号码的字母组合
https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
*/
func letterCombinations(digits string) []string {

	phoneMap := map[int][]string{}
	phoneMap[2] = []string{"a", "b", "c"}
	phoneMap[3] = []string{"d", "e", "f"}
	phoneMap[4] = []string{"g", "h", "i"}
	phoneMap[5] = []string{"j", "k", "l"}
	phoneMap[6] = []string{"m", "n", "o"}
	phoneMap[7] = []string{"p", "q", "r", "s"}
	phoneMap[8] = []string{"t", "u", "v"}
	phoneMap[9] = []string{"w", "x", "y", "z"}

	combinationMap := map[string]int{}

	for _, i2 := range digits {

		num := i2 - '0'
		strings := phoneMap[int(num)]

		if len(combinationMap) == 0 {

			for _, s := range strings {
				combinationMap[s] = 0
			}

		} else {
			tempMap := map[string]int{}
			for key := range combinationMap {
				for _, item := range strings {
					temp := key + item
					tempMap[temp] = 0
				}
			}
			combinationMap = tempMap

		}

	}

	result := make([]string, len(combinationMap))

	index := 0
	for key, _ := range combinationMap {
		result[index] = key
		index++
	}

	return result
}

func main() {

	digits := "2"

	fmt.Printf("%v\n", letterCombinations(digits))

}
