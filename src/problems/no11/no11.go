package main

import (
	"fmt"
	"math"
)

/*
11. 盛最多水的容器
https://leetcode.cn/problems/container-with-most-water/
*/
func maxArea(height []int) int {

	length := len(height)

	max := 0

	var increaseIndex []int

	increaseIndex = append(increaseIndex, 0)

	for i := 1; i < length; i++ {

		temp := height[i]

		for j := 0; j < len(increaseIndex); j++ {

			tempHeight := int(math.Min(float64(height[increaseIndex[j]]), float64(temp)))
			tempLength := i - increaseIndex[j]
			max = int(math.Max(float64(max), float64(tempHeight*tempLength)))
		}

		if temp > height[increaseIndex[len(increaseIndex)-1]] {
			increaseIndex = append(increaseIndex, i)
		}
	}

	return max
}

func main() {

	height := []int{1, 1}

	fmt.Println(maxArea(height))

}
