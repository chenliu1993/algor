package main

import (
	"fmt"
)

func threeSum(num int) [][]int {
	var (
		ans     [][]int
		visited []int
		cur     []int
		iterate func(int, int)
	)
	ans = [][]int{}
	cur = []int{}
	visited = make([]int, num)
	iterate = func(sum, counter int) {
		if sum > num {
			return
		}
		if sum == num {
			if counter == 2 {
				temp := make([]int, counter)
				copy(temp, cur)
				ans = append(ans, temp)
			}
			return
		}
		for i := 0; i < num; i++ {
			if visited[i] != 1 {
				if len(cur) != 0 {
					if cur[len(cur)-1] > i {
						continue
					}
				}
				visited[i] = 1
				cur = append(cur, i)
				sum = sum + i
				iterate(sum, counter+1)
				sum = sum - i
				cur = cur[:len(cur)-1]
				visited[i] = 0
			}
		}
	}
	iterate(0, 0)
	return ans
}

func main() {
	num := 6
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	fmt.Println(threeSum(num))
}
