/*
 * @lc app=leetcode.cn id=659 lang=golang
 *
 * [659] 分割数组为连续子序列
 */

// @lc code=start
func isPossible(nums []int) bool {
	n := len(nums)
	var (
		resultStk []int
		stk       []int
		dict      []int
		record    map[int]int
	)
	dict = []int{}
	record = map[int]int{}
	for i := 0; i < n; i++ {
		if _, ok := record[nums[i]]; !ok {
			record[nums[i]] = 1
			dict = append(dict, nums[i])
		} else {
			record[nums[i]]++
		}
	}
	dictLen := len(dict)
	for dictLen != 0 {
		lastV := record[dict[0]]
		for i := 0; i < len(dict); i++ {
			if len(stk) != 0 {
				if dict[i]-stk[len(stk)-1] > 1 || record[dict[i]] < lastV {
					if len(stk) < 3 {
						return false
					}
					break
				}
			}
			stk = append(stk, dict[i])
			lastV = record[dict[i]]
			record[dict[i]] = record[dict[i]] - 1
		}
		tempDict := []int{}
		for i := 0; i < len(dict); i++ {
			if record[dict[i]] != 0 {
				tempDict = append(tempDict, dict[i])
			} else {
				delete(record, dict[i])
			}
		}

		dictLen = len(tempDict)
		fmt.Println(dict, record, dictLen)
		dict = tempDict
		resultStk = make([]int, len(stk))
		copy(resultStk, stk)
		stk = []int{}
	}
	return len(resultStk) >= 3
}

// @lc code=end

