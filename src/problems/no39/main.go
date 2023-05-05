package main

import "fmt"

var combinations [][]int

/*
*
39. 组合总和
https://leetcode.cn/problems/combination-sum/
*/
func combinationSum(candidates []int, target int) [][]int {

	combinations = [][]int{}

	doCombinationSum(candidates, []int{}, target)

	return combinations
}

func doCombinationSum(candidates []int, preIndex []int, target int) {

	sum := 0
	for i := 0; i < len(preIndex); i++ {
		sum += candidates[preIndex[i]]
	}

	if sum > target {
		return
	}

	if sum == target {

		var temp []int
		for i := 0; i < len(preIndex); i++ {
			temp = append(temp, candidates[preIndex[i]])
		}

		combinations = append(combinations, temp)
		return
	}

	index := 0
	if len(preIndex) != 0 {
		index = preIndex[len(preIndex)-1]
	}

	for index < len(candidates) {

		curIndex := append(preIndex, index)

		doCombinationSum(candidates, curIndex, target)

		index++
	}

}

func main() {

	candidates := []int{2, 3, 6, 7}

	target := 7

	sum := combinationSum(candidates, target)

	fmt.Printf("%v\n", sum)

}
