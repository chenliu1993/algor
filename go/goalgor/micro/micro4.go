package main

import (
	"fmt"
)

func sqrt(input int64) (result int64) {
	limit := int64(1)
	result = 1
	// Specify the case for input is 0
	if input == 0 {
		result = 0
	}
	for limit = input/2 + 1; limit >= 1; limit = limit / 2 {
		if limit*limit-input > 0 && (limit-1)*(limit-1)-input <= 0 {
			result = limit - 1
		}
	}
	return result
}

func main() {
	var length = 16
	var inputs = make([]int64, int64(length))
	for i := 0; i < length; i++ {
		inputs[i] = int64(i)
	}
	fmt.Printf("\n/****output****/\n")
	for i := 0; i < length; i++ {
		fmt.Printf("position %d is %d\n", i, sqrt(inputs[i]))
	}
}
