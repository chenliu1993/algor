package main

import (
	"fmt"
	"strconv"
)

func add(num1 string, num2 string) string {
	n1 := len(num1)
	n2 := len(num2)
	carry := 0
	sum := 0
	i := n1 - 1
	j := n2 - 1
	var ans string
	for i >= 0 && j >= 0 {
		num1Int, err := strconv.Atoi(string(num1[i]))
		if err != nil {
			panic(err)
		}

		num2Int, err := strconv.Atoi(string(num2[j]))
		if err != nil {
			panic(err)
		}
		sum = num1Int + num2Int + carry
		sum, carry = sum%10, sum/10
		ans = strconv.Itoa(sum) + ans
		i = i - 1
		j = j - 1
	}
	for i >= 0 {
		sum = 0
		num1Int, err := strconv.Atoi(string(num1[i]))
		if err != nil {
			panic(err)
		}
		sum = num1Int + carry
		sum, carry = sum%10, sum/10
		ans = strconv.Itoa(sum) + ans
		i = i - 1
	}
	for j >= 0 {
		sum = 0
		num2Int, err := strconv.Atoi(string(num2[j]))
		if err != nil {
			panic(err)
		}
		sum = num2Int + carry
		sum, carry = sum%10, sum/10
		ans = strconv.Itoa(sum) + ans
		j = j - 1
	}
	if carry != 0 {
		ans = strconv.Itoa(carry) + ans
	}
	return ans
}

func multiply(num1 string, num2 string) string {
	n1 := len(num1)
	n2 := len(num2)
	// Make sure n1 is the longest one
	if n1 < n2 {
		temp := num1
		num1 = num2
		num2 = temp
	}
	n1 = len(num1)
	n2 = len(num2)

	totalSum := "0"
	count := 0
	for i := n2 - 1; i >= 0; i-- {
		num2Int, err := strconv.Atoi(string(num2[i]))
		if err != nil {
			panic(err)
		}
		sum := "0"
		for j := 0; j < num2Int; j++ {
			sum = add(sum, num1)
		}
		for k := 0; k < count; k++ {
			sum = sum + "0"
		}
		totalSum = add(totalSum, sum)
		count = count + 1
	}
	return totalSum
}

func main() {
	num1 := "0"
	num2 := "52"

	fmt.Println(multiply(num1, num2))
}
