package no39

func combinationSum(candidates []int, target int) [][]int {

	result = make([][]int, 0)

	dfs(nil, candidates, 0, target)

	return result
}

var result [][]int

func dfs(preNums, resNums []int, preSum, target int) {

	if preSum == target {
		result = append(result, preNums)
	} else if preSum > target {
		return
	}

	for i, num := range resNums {

		curNums := make([]int, 0)
		curNums = append(curNums, preNums...)
		curNums = append(curNums, num)

		curResNums := make([]int, 0)
		curResNums = append(curResNums, resNums[i:]...)

		dfs(curNums, curResNums, preSum+num, target)
	}

}
