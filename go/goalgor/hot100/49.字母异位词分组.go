/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{strs}
	}
	res := [][]string{}
	temp := map[string][]string{}
	for i := 0; i < len(strs); i++ {
		strSlice := strings.Split(strs[i], "")
		sort.Strings(strSlice)
		newstr := strings.Join(strSlice, "")
		temp[newstr] = append(temp[newstr], strs[i])
	}
	for _, v := range temp {
		res = append(res, v)
	}
	return res
}

// @lc code=end

