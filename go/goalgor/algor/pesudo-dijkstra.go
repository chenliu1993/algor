package algor

func minD(src int, adjacentRelation [][]int, numNode int, visited []int) []int {
	var (
		countQueue []int
		rear       int
		front      int
		// minDistance needs to be initialized with a +inf
		minDistance []int
	)
	rear = 0
	front = 0
	countQueue = make([]int, numNode)
	minDistance = make([]int, numNode)
	countQueue[rear] = src
	rear++
	minDistance[src] = 0
	for {
		if rear == front {
			break
		}
		for start := front; start < rear; start++ {
			if visited[start] != 0 {
				continue
			}
			tempRear := rear
			for i := 0; adjacentRelation[start][i] != 0 && i < numNode; i++ {
				if i == start {
					continue
				}
				// start to i has a direct road.
				if adjacentRelation[start][i] != -1 {
					countQueue[tempRear] = i
					tempRear++
				}
				if minDistance[start]+adjacentRelation[start][i] < minDistance[i] {
					minDistance[i] = minDistance[start] + adjacentRelation[start][i]
				}
			}
			visited[start] = 1
		}
		front = rear
	}
	return minDistance
}
