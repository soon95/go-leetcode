package main

import "fmt"

/*
42. 接雨水
*https://leetcode.cn/problems/trapping-rain-water/
*/
func trap(height []int) int {

	length := len(height)

	dp := make([]int, length)

	dp[0] = 0

	for i := 1; i < length; i++ {

		cur := height[i]

		// 向前寻找开始位置
		start := i - 1
		for j := i - 1; j >= 0; j-- {

			if cur <= height[j] {
				// 如果找到比当前高的，找到了
				start = j
				break
			} else {
				// 要不然，记录一下到目前为止找到的最高的
				if height[j] > height[start] {
					start = j
				}
			}
		}

		// 确认一下目前的最大高度
		target := 0
		if height[start] > cur {
			target = cur
		} else {
			target = height[start]
		}

		temp := 0
		for k := start + 1; k < i; k++ {
			temp += target - height[k]
		}

		dp[i] = dp[start] + temp
	}

	return dp[length-1]
}

func main() {

	height := []int{4, 2, 0, 3, 2, 5}

	fmt.Println(trap(height))

}
