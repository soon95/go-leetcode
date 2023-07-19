package main

import "fmt"

/*
*https://leetcode.cn/problems/course-schedule/
207. 课程表
通过 入度来判断
*/
func canFinish(numCourses int, prerequisites [][]int) bool {

	dic := make([][]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		dic[i] = make([]bool, numCourses)
	}

	inTimes := make([]int, numCourses)

	for _, prerequisite := range prerequisites {

		from := prerequisite[1]
		to := prerequisite[0]

		dic[from][to] = true

		inTimes[to] = inTimes[to] + 1
	}

	temp := []int{}

	for num, time := range inTimes {
		if time == 0 {
			temp = append(temp, num)
		}
	}

	for len(temp) != 0 {

		nextTemp := []int{}

		for _, start := range temp {

			for end, connect := range dic[start] {

				if connect {
					inTimes[end] = inTimes[end] - 1

					if inTimes[end] == 0 {
						nextTemp = append(nextTemp, end)
					}
				}

			}
		}

		temp = nextTemp
	}

	for _, time := range inTimes {
		if time != 0 {
			return false
		}
	}

	return true
}

func main() {

	//numCourses := 3
	//prerequisites := [][]int{
	//	{1, 0}, {0, 2}, {2, 1},
	//}
	//numCourses := 2
	//prerequisites := [][]int{
	//	{1, 0},
	//}
	numCourses := 2
	prerequisites := [][]int{
		{1, 0}, {0, 1},
	}
	//numCourses := 5
	//prerequisites := [][]int{
	//	{1, 4}, {2, 4}, {3, 1}, {3, 2},
	//}

	fmt.Printf("%v", canFinish(numCourses, prerequisites))
}
