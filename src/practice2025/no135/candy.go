package no135

func candy(ratings []int) int {

	tmp := make([]int, len(ratings))

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			tmp[i] = tmp[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			tmp[i] = max(tmp[i], tmp[i+1]+1)
		}
	}

	cnt := len(ratings)
	for _, num := range tmp {
		cnt += num
	}
	return cnt
}
