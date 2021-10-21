/*
 * @lc app=leetcode.cn id=319 lang=golang
 *
 * [319] 灯泡开关
 */

// @lc code=start
// func bulbSwitch(n int) int {
// 	lights := make([]int, n)
// 	for i := 1; i <= n; i++ {
// 		for j := i - 1; j < n; j = j + i {
// 			if lights[j] == 1 {
// 				lights[j] = 0
// 			} else {
// 				lights[j] = 1
// 			}
// 		}
// 	}
// 	sum := 0
// 	for i := 0; i < len(lights); i++ {
// 		sum = sum + lights[i]
// 	}
// 	return sum
// }

func bulbSwitch(n int) int {
	count := 0
	for i := 1; i*i <= n; i++ {
		count = count + 1
	}
	return count
}

// @lc code=end

