package main

import (
	"fmt"
	"strconv"
)

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func flipLights(n int, presses int) int {
	var (
		curSum int
		curOp  []int
		ans    map[string]int
		ops    [][]int
		getOps func(int)
	)
	getOps = func(opcount int) {
		if opcount == 4 {
			if curSum == presses {
				tempOp := make([]int, len(curOp))
				copy(tempOp, curOp)
				ops = append(ops, tempOp)
			}
			return
		}
		for i := 0; i <= presses-curSum; i++ {
			curOp = append(curOp, i%2)
			curSum = curSum + i
			getOps(opcount + 1)
			curOp = curOp[:len(curOp)-1]
			curSum = curSum - i
		}
	}
	curSum = 0
	canChange := 0
	curOp = []int{}
	ops = [][]int{}
	ans = map[string]int{}
	getOps(0)
	for i := 0; i < len(ops); i++ {
		canChange = 0
		tempRes := ""
		for j := 1; j <= Min(3, n); j++ {
			if j%2 == 0 {
				canChange = ops[i][0] + ops[i][1]
			} else {
				canChange = ops[i][0] + ops[i][2]
			}
			if j%3 == 1 {
				canChange = canChange + ops[i][3]
			}
			tempRes = tempRes + strconv.Itoa(canChange%2)
		}
		if _, ok := ans[tempRes]; !ok {
			ans[tempRes] = 1
		}
	}
	return len(ans)
}

func main() {
	n := 5
	presses := 1000
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println(flipLights(n, presses))
}

// But over time limit
