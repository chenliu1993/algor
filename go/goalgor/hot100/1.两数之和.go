/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start

func twoSum(nums []int, target int) []int {
	ans := []int{}
	// O(n^2)
	// for i := 0; i < len(nums)-1; i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[i]+nums[j] == target {
	// 			ans = append(ans, i, j)
	// 		}
	// 	}
	// }
	record := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		record[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		if _, ok := record[target-nums[i]]; ok && record[target-nums[i]] != i {
			ans = append(ans, i, record[target-nums[i]])
			break
		}
	}
	return ans
}

// @lc code=end

