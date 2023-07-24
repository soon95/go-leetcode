package main

import (
	"fmt"
)

/*
*621. 任务调度器
https://leetcode.cn/problems/task-scheduler/
*/
func leastInterval(tasks []byte, n int) int {

	maxTaskNum := 0
	dic := map[byte]int{}
	for _, task := range tasks {
		dic[task]++

		if maxTaskNum < dic[task] {
			maxTaskNum = dic[task]
		}
	}

	rest := 0
	for _, num := range dic {
		if num == maxTaskNum {
			rest++
		}
	}

	temp := (maxTaskNum-1)*(n+1) + rest

	totalTaskNums := len(tasks)

	// 不可能小于totalTaskNums
	if temp <= totalTaskNums {
		return totalTaskNums
	} else {
		return temp
	}
}

func main() {

	//tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	//n := 2

	//tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	//n := 0

	//tasks := []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	//n := 1

	tasks := []byte{'A', 'B', 'C', 'D', 'A', 'B', 'V'}
	n := 3

	fmt.Printf("%v", leastInterval(tasks, n))

}
