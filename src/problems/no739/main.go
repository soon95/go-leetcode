package main

import "fmt"

/*
*739. 每日温度
https://leetcode.cn/problems/daily-temperatures/
*/
func dailyTemperatures(temperatures []int) []int {

	length := len(temperatures)

	ans := make([]int, length)

	// 维护一个单调栈

	indexStack := []int{}

	indexStack = append(indexStack, 0)

	for i := 1; i < length; i++ {

		for len(indexStack) != 0 && temperatures[i] > temperatures[indexStack[len(indexStack)-1]] {
			// 如果当前天气大于栈顶天气

			topIndex := indexStack[len(indexStack)-1]

			ans[topIndex] = i - topIndex

			indexStack = indexStack[0 : len(indexStack)-1]
		}

		indexStack = append(indexStack, i)
	}

	return ans
}

func main() {

	//temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	temperatures := []int{30, 40, 50, 60}

	fmt.Printf("%v", dailyTemperatures(temperatures))

}
