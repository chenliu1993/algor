/*
 * @lc app=leetcode.cn id=721 lang=golang
 *
 * [721] 账户合并
 */

// @lc code=start
func accountsMerge(accounts [][]string) [][]string {
	var (
		fa            []int
		emailToName   map[string]string
		indexToEmails map[int][]string
		emailToIndex  map[string]int
		find          func(int) int
		merge         func(int, int)
	)
	emailToName = map[string]string{}
	emailToIndex = map[string]int{}
	for i := 0; i < len(accounts); i++ {
		for j := 1; j < len(accounts[i]); j++ {
			if _, ok := emailToIndex[accounts[i][j]]; !ok {
				emailToName[accounts[i][j]] = accounts[i][0]
				emailToIndex[accounts[i][j]] = len(emailToIndex)
			}
		}
	}
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge = func(x, y int) {
		fa[find(x)] = find(y)
	}
	fa = make([]int, len(emailToIndex))
	for i := 0; i < len(fa); i++ {
		fa[i] = i
	}
	for i := 0; i < len(accounts); i++ {
		firstIndex := emailToIndex[accounts[i][1]]
		for j := 2; j < len(accounts[i]); j++ {
			merge(emailToIndex[accounts[i][j]], firstIndex)
		}
	}
	indexToEmails = map[int][]string{}
	for k, v := range emailToIndex {
		v = find(v)
		indexToEmails[v] = append(indexToEmails[v], k)
	}
	ans := [][]string{}
	for _, emails := range indexToEmails {
		sort.Strings(emails)
		account := append([]string{emailToName[emails[0]]}, emails...)
		ans = append(ans, account)
	}
	return ans
}

// @lc code=end

