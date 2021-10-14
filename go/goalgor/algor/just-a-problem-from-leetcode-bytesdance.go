package main

import (
	"fmt"
	"log"
	"strconv"
)

func findLongestStr(str string) string {
	var (
		first  int
		rear   int
		ans    int
		maxStr string
	)
	maxStr = ""
	for rear != len(str) {
		// Get first number
		for rear < len(str) && (str[rear] < '1' || str[rear] > '9') {
			rear = rear + 1
		}
		first = rear
		// Grt first non-number
		for rear < len(str) && (str[rear] >= '1' && str[rear] <= '9') {
			rear = rear + 1
		}
		cur, err := strconv.Atoi(str[first:rear])
		if err != nil {
			log.Println(err)
			return ""
		}
		// Compare to get the result
		if cur > ans {
			ans = cur
			maxStr = str[first:rear]
		}
		// Conintue to next loop
		first = rear
	}
	return maxStr
}

func main() {
	str := "sif123sdqw342xxcs54321sx12345"
	fmt.Println(findLongestStr(str))
}
