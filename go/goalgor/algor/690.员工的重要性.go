/*
 * @lc app=leetcode.cn id=690 lang=golang
 *
 * [690] 员工的重要性
 */

// @lc code=start
/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

func getImportance(employees []*Employee, id int) int {
	var (
		record, idToIdx       map[int]int
		getEmployeeImportance func(int) int
	)
	record = map[int]int{}
	idToIdx = map[int]int{}
	for idx, val := range employees {
		record[val.Id-1] = -1
		idToIdx[val.Id] = idx
	}
	getEmployeeImportance = func(start int) int {
		if len(employees[start].Subordinates) == 0 {
			record[employees[start].Id-1] = employees[start].Importance
			return employees[start].Importance
		}
		cur := employees[start].Importance
		for _, val := range employees[start].Subordinates {
			if record[val-1] != -1 {
				cur += record[val-1]
			} else {
				cur += getEmployeeImportance(idToIdx[val])
			}
		}
		record[employees[start].Id-1] = cur
		return cur
	}
	for k, v := range idToIdx {
		record[k-1] = getEmployeeImportance(v)
	}
	return record[id-1]
}

// @lc code=end

// Just to get id
// func getImportance(employees []*Employee, id int) (total int) {
// 	mp := map[int]*Employee{}
// 	for _, employee := range employees {
// 		mp[employee.Id] = employee
// 	}

// 	var dfs func(int)
// 	dfs = func(id int) {
// 		employee := mp[id]
// 		total += employee.Importance
// 		for _, subId := range employee.Subordinates {
// 			dfs(subId)
// 		}
// 	}
// 	dfs(id)
// 	return
// }

