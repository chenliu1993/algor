package main

import (
	"fmt"
)

// var dict = map[int]string{
// 	1000: "M",
// 	900:  "CM",
// 	500:  "D",
// 	400:  "CD",
// 	100:  "C",
// 	90:   "XC",
// 	50:   "L",
// 	40:   "XL",
// 	10:   "X",
// 	9:    "IX",
// 	5:    "V",
// 	4:    "IV",
// 	1:    "I",
// }

// func getNum(n int) string {
// 	if n < 10 {
// 		if n == 4 || n == 9 {
// 			return dict[n]
// 		}
// 		if n < 5 {
// 			return strings.Repeat(dict[1], n)
// 		} else if n < 10 && n >= 5 {
// 			return dict[5] + strings.Repeat(dict[1], n-5)
// 		}
// 	} else if n >= 10 && n < 100 {
// 		if n == 40 || n == 90 || n == 10 || n == 50 {
// 			return dict[n]
// 		}
// 		if n < 50 {
// 			return strings.Repeat(dict[10], n/10)
// 		} else if n > 50 && n < 100 {
// 			return dict[50] + strings.Repeat(dict[10], (n-50)/10)
// 		}
// 	} else if n >= 100 && n < 1000 {
// 		if n == 100 || n == 400 || n == 900 || n == 500 {
// 			return dict[n]
// 		}
// 		if n < 500 {
// 			return strings.Repeat(dict[100], n/100)
// 		} else if n > 500 && n < 1000 {
// 			return dict[500] + strings.Repeat(dict[100], (n-500)/100)
// 		}
// 	}
// 	if n == 1000 {
// 		return dict[n]
// 	}
// 	return strings.Repeat(dict[1000], n/1000)
// }

// func intToRoman(num int) string {
// 	ans := ""
// 	bit := 1
// 	var carry int
// 	for num != 0 {
// 		carry = (num % 10) * bit
// 		ans = getNum(carry) + ans
// 		num = num / 10
// 		bit = bit * 10
// 	}
// 	return ans
// }

var valueSymbols = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	ans := ""
	for _, v := range valueSymbols {
		for num >= v.value {
			num = num - v.value
			ans += v.symbol
		}
		if num == 0 {
			break
		}
	}
	return ans
}

func main() {
	num := 1994
	fmt.Println(intToRoman(num))
}
