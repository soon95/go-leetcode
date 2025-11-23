package no202

func isHappy(n int) bool {

	tmp := make(map[int]bool)

	for {

		n = transform(n)

		if n == 1 {
			return true
		}

		if _, exist := tmp[n]; exist {
			return false
		}

		tmp[n] = true
	}
}

func transform(n int) int {

	ans := 0

	for n > 0 {

		low := n % 10
		high := n / 10

		ans += low * low

		n = high
	}

	return ans
}
