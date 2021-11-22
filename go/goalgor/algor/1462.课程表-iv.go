/*
 * @lc app=leetcode.cn id=1462 lang=golang
 *
 * [1462] 课程表 IV
 */

// @lc code=start
func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	var (
		record [][]bool
	)
	record = make([][]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		record[i] = make([]bool, numCourses)
	}
	for i := 0; i < len(prerequisites); i++ {
		record[prerequisites[i][0]][prerequisites[i][1]] = true
	}

	for mid := 0; mid < numCourses; mid++ {
		for i := 0; i < numCourses; i++ {
			for j := 0; j < numCourses; j++ {
				record[i][j] = record[i][j] || (record[i][mid] && record[mid][j])
			}
		}
	}
	ans := []bool{}
	for i := 0; i < len(queries); i++ {
		if record[queries[i][0]][queries[i][1]] {
			ans = append(ans, true)
			continue
		}
		ans = append(ans, false)
	}
	return ans
}

// @lc code=end

