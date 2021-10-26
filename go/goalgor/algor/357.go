package main

import (
	"fmt"
	"log"
)

func Pow10(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	if n%2 == 0 {
		return Pow10(n/2) * Pow10(n/2)
	}
	return Pow10((n-1)/2) * Pow10((n-1)/2) * 10
}

func countNumbersWithUniqueDigits(n int) int {
	var (
		count     int
		rightMost int = Pow10(n)
		check     func(int) bool
	)
	check = func(x int) bool {
		dict := map[int]bool{}
		for x != 0 {
			left := x % 10
			x = x / 10
			if _, ok := dict[left]; ok {
				return false
			} else {
				dict[left] = true
			}
		}
		return true
	}
	// 0 is the default one
	count = 1
	for i := 1; i < rightMost; i++ {
		if check(i) {
			count = count + 1
		}
	}
	return count
}

func main() {
	n := 8
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	fmt.Println(countNumbersWithUniqueDigits(n))
}

// Over time limit
