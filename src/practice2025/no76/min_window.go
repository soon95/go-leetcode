package no76

func minWindow(s string, t string) string {

	sourceAlphaMap, targetAlphaMap := make(map[int32]int), make(map[int32]int)

	for _, ch := range t {
		targetAlphaMap[ch] = targetAlphaMap[ch] + 1
	}

	left, right := 0, 0
	sourceAlphaMap[int32(s[0])] = 1
	res := ""
	moveRight := true
	for (right < len(s) || !moveRight) && left <= right {

		if alphaContainsAll(sourceAlphaMap, targetAlphaMap) {
			// 如果找到了 那就收缩
			if res == "" || len(res) > right-left+1 {
				res = s[left : right+1]
			}
			moveRight = false
		} else {
			moveRight = true
		}

		if moveRight {
			right++

			if right == len(s) {
				break
			}

			sourceAlphaMap[int32(s[right])] = sourceAlphaMap[int32(s[right])] + 1

		} else {
			sourceAlphaMap[int32(s[left])] = sourceAlphaMap[int32(s[left])] - 1

			left++
		}

	}
	return res
}

func alphaContainsAll(source, target map[int32]int) bool {

	for alpha, cnt := range target {

		if source[alpha] < cnt {
			return false
		}

	}
	return true
}
