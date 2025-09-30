package no78

func subsets(nums []int) [][]int {
	return dfs(nil, nums)
}

func dfs(preNums, resNums []int) [][]int {

	res := [][]int{preNums}

	for i, num := range resNums {

		curNums := make([]int, 0)
		curNums = append(curNums, preNums...)
		curNums = append(curNums, num)

		curResNums := make([]int, 0)
		curResNums = append(curResNums, resNums[i+1:]...)

		res = append(res, dfs(curNums, curResNums)...)
	}

	return res
}
