package no207

func canFinish(numCourses int, prerequisites [][]int) bool {

	// 节点入度
	inMap := make(map[int]map[int]bool)
	for i := 0; i < numCourses; i++ {
		inMap[i] = make(map[int]bool)
	}
	outMap := make(map[int][]int)

	for _, prerequisite := range prerequisites {
		inMap[prerequisite[0]][prerequisite[1]] = true
		outMap[prerequisite[1]] = append(outMap[prerequisite[1]], prerequisite[0])
	}

	visited := make(map[int]bool)

	visitedCnt := len(visited)
	for {

		for i, preMap := range inMap {
			if len(preMap) == 0 {
				// 可以访问
				visited[i] = true
				for _, post := range outMap[i] {
					delete(inMap[post], i)
				}
			}
		}

		if visitedCnt == len(visited) {
			break
		}

		visitedCnt = len(visited)
	}

	return visitedCnt == numCourses
}
