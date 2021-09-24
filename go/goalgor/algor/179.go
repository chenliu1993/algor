// package main

// import (
// 	"fmt"
// 	"math"
// 	"strconv"
// 	"strings"
// )

// func sortCombine(nums []int) string {
// 	results := [][]string{}
// 	result := []string{}
// 	n := len(nums)
// 	visited := make([]int, n)
// 	count := 0
// 	var combine func()
// 	combine = func() {
// 		if count >= n {
// 			tempResult := []string{}
// 			for i := 0; i < len(result); i++ {
// 				tempResult = append(tempResult, result[i])
// 			}
// 			results = append(results, tempResult)
// 			return
// 		}
// 		for i := 0; i < n; i++ {
// 			if visited[i] != 1 {
// 				result = append(result, strconv.Itoa(nums[i]))
// 				visited[i] = 1
// 				count = count + 1
// 				combine()
// 				result = result[:len(result)-1]
// 				visited[i] = 0
// 				count = count - 1
// 			}
// 		}
// 	}
// 	combine()
// 	var (
// 		sum     uint64
// 		maxi    int
// 		maxSum  uint64
// 		lastStr string
// 	)
// 	maxSum = uint64(0)
// 	fmt.Println(results)
// 	for i := 0; i < len(results); i++ {
// 		sum = uint64(0)
// 		lastStr = ""
// 		for j := len(results[i]) - 1; j >= 0; j-- {
// 			res, err := strconv.ParseUint(results[i][j], 10, 64)
// 			if err != nil {
// 				panic(err)
// 			}
// 			fmt.Println(res * uint64(math.Pow10(len(lastStr))))
// 			sum = sum + res*uint64(math.Pow10(len(lastStr)))
// 			lastStr = lastStr + results[i][j]
// 		}
// 		if maxSum < sum {
// 			maxSum = sum
// 			maxi = i
// 		}
// 	}
// 	return strings.Join(results[maxi], "")
// }

// func largestNumber(nums []int) string {
// 	return sortCombine(nums)
// }

// func main() {
// 	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
// 	fmt.Println(largestNumber(nums))
// }

// Above is out of space

// package main

// import (
// 	"fmt"
// 	"math"
// 	"strconv"
// )

// func sortCombine(nums []int) string {
// 	result := []string{}
// 	n := len(nums)
// 	visited := make([]int, n)
// 	count := 0
// 	var (
// 		sum        uint64
// 		maxSum     uint64
// 		lastStr    string
// 		combine    func()
// 		tempResult []string
// 	)
// 	maxSum = uint64(0)
// 	combine = func() {
// 		if count >= n {
// 			tempResult = make([]string, len(result))
// 			copy(tempResult, result)
// 			sum = uint64(0)
// 			lastStr = ""
// 			for j := len(result) - 1; j >= 0; j-- {
// 				res, err := strconv.ParseUint(result[j], 10, 64)
// 				if err != nil {
// 					panic(err)
// 				}
// 				sum = sum + res*uint64(math.Pow10(len(lastStr)))
// 				lastStr = lastStr + result[j]
// 			}
// 			if maxSum < sum {
// 				maxSum = sum
// 			}
// 			return
// 		}
// 		for i := 0; i < n; i++ {
// 			if visited[i] != 1 {
// 				result = append(result, strconv.Itoa(nums[i]))
// 				visited[i] = 1
// 				count = count + 1
// 				combine()
// 				result = result[:len(result)-1]
// 				visited[i] = 0
// 				count = count - 1
// 			}
// 		}
// 	}
// 	combine()
// 	return strconv.FormatUint(maxSum, 10)
// }

// func largestNumber(nums []int) string {
// 	return sortCombine(nums)
// }

// func main() {
// 	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
// 	fmt.Println(largestNumber(nums))
// }

// Above is out of time

// Convert to sort

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x := nums[i]
		y := nums[j]
		sx := 10
		sy := 10
		for sx <= x {
			sx = sx * 10
		}
		for sy <= y {
			sy = sy * 10
		}
		return sx*y+x < sy*x+y
	})
	if nums[0] == 0 {
		return "0"
	}
	res := []byte{}
	for i := 0; i < len(nums); i++ {
		res = append(res, strconv.Itoa(nums[i])...)
	}
	return string(res)
}