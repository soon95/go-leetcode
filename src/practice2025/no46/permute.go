package no46

func permute(nums []int) [][]int {
	return dfs(nil, nums)
}

func dfs(preNums, resNums []int) [][]int {

	if len(resNums) == 0 {
		return [][]int{preNums}
	}

	res := make([][]int, 0)

	for i, num := range resNums {

		nextPre := make([]int, 0)
		nextPre = append(nextPre, preNums...)
		nextPre = append(nextPre, num)

		nextRes := make([]int, 0)
		nextRes = append(nextRes, resNums[:i]...)
		nextRes = append(nextRes, resNums[i+1:]...)

		res = append(res, dfs(nextPre, nextRes)...)
	}

	return res
}
