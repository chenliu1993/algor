/*
 * @lc app=leetcode.cn id=813 lang=golang
 *
 * [813] 最大平均值和的分组
 */

// @lc code=start
func largestSumOfAverages(nums []int, k int) float64 {
	n := len(nums)
	var (
		sum      int
		curMax   float64
		subSlice int
		record   [][]float64
	)
	record = make([][]float64, n)
	for i := 0; i < n; i++ {
		record[i] = make([]float64, k)
	}
	sum = 0
	for i := 0; i < n; i++ {
		sum = sum + nums[i]
		record[i][0] = float64(sum) / float64(i+1)
	}
	for i := 0; i < n; i++ {
		for j := 1; j < k; j++ {
			subSlice = 0
			curMax = float64(0)
			for m := 0; m < i; m++ {
				subSlice = subSlice + nums[i-m]
				curMax = math.Max(curMax, record[i-m-1][j-1]+float64(subSlice)/float64(m+1))
			}
			record[i][j] = curMax
		}
	}
	fmt.Println(record)
	return record[n-1][k-1]
}

// @lc code=end

