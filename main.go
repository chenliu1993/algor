package main

import (
	"fmt"
	"strconv"
	"strings"
)

func maximumSwap(num int) int {
	var (
		j     int
		ans   int
		str   []string
		first bool
		stk   []int
	)
	first = true
	str = strings.Split(strconv.Itoa(num), "")
	stk = []int{}
	for j = 0; j < len(str); j++ {
		for len(stk) != 0 && str[stk[len(stk)-1]] < str[j] {
			first = false
			stk = stk[:len(stk)-1]
		}
		if !first {
			break
		}
		stk = append(stk, j)

	}
	if j == len(str) {
		return num
	}
	spliter := j
	maxEle := str[j]
	max := j
	for ; j < len(str); j++ {
		fmt.Println(str[j])
		if str[j] >= maxEle {
			max = j
			maxEle = str[j]
		}
	}
	fmt.Println(max, str[max])
	min := 0
	for i := 0; i < spliter; i++ {
		if str[i] < maxEle {
			min = i
			break
		}
	}
	str[min], str[max] = str[max], str[min]
	ans, _ = strconv.Atoi(strings.Join(str, ""))
	return ans
}

func main() {
	num := 12335431
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	fmt.Println(maximumSwap(num))
}
