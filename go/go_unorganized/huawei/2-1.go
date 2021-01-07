package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	val   string
	left  *TreeNode
	right *TreeNode
}

func calculate(s string, a, b int) int {
	if strings.Compare(s, "add") == 0 {
		return a + b
	}
	if strings.Compare(s, "sub") == 0 {
		return a - b
	}
	if strings.Compare(s, "mul") == 0 {
		return a * b
	}
	if strings.Compare(s, "div") == 0 {
		if b == 0 {
			fmt.Println("error")
			panic(0)
		}
	}
	return a / b
}

func check(s string) bool {
	if strings.Compare(s, "add") == 0 {
		return true
	}
	if strings.Compare(s, "sub") == 0 {
		return true
	}
	if strings.Compare(s, "mul") == 0 {
		return true
	}
	if strings.Compare(s, "div") == 0 {
		return true
	}
	return false
}

func calSum(inputs []string) int {
	stk := []string{}
	rear := 0
	for idx := range inputs {
		if check(inputs[idx]) {
			stk = append(stk, inputs[idx])
			rear++
			continue
		}
		if rear != 0 && check(stk[rear-1]) {
			stk = append(stk, inputs[idx])
			rear++
			continue
		}
		p1 := stk[rear-1]
		rear--
		op := stk[rear-1]
		p2 := inputs[idx]
		p1Val, _ := strconv.Atoi(p1)
		p2Val, _ := strconv.Atoi(p2)
		stk[rear-1] = strconv.Itoa(calculate(op, p1Val, p2Val))
		fmt.Println(stk)
	}
	res, _ := strconv.Atoi(stk[0])
	return res
}

func main() {
	var input string
	input = "(sub 2 mul 2 4)"
	vals := strings.Split(input, " ")
	for idx := range vals {
		if strings.Compare(string(vals[idx][0]), "(") == 0 {
			vals[idx] = vals[idx][1:len(vals[idx])]
			continue
		}
		if strings.Compare(string(vals[idx][len(vals[idx])-1]), ")") == 0 {
			vals[idx] = vals[idx][:len(vals[idx])-1]
			continue
		}
	}
	fmt.Println(vals)
	total := calSum(vals)
	fmt.Println(total)
	return
}
