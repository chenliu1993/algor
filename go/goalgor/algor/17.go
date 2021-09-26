package main

import (
	"fmt"
	"strings"
)

var dict = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	n := len(digits)
	if n == 0 {
		return []string{}
	}
	count := 0
	results := []string{}
	result := []string{}
	var iterate func(int)
	iterate = func(start int) {
		if count == n {
			tempResult := make([]string, len(result))
			copy(tempResult, result)
			results = append(results, strings.Join(tempResult, ""))
			return
		}
		for i := start; i < n; i++ {
			cur := dict[string(digits[i])]
			for j := 0; j < len(cur); j++ {
				result = append(result, string(cur[j]))
				count = count + 1
				iterate(i + 1)
				count = count - 1
				result = result[:len(result)-1]
			}
		}
	}
	iterate(0)
	return results
}

func main() {
	fmt.Println(letterCombinations("2"))
}
