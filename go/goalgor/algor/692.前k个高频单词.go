/*
 * @lc app=leetcode.cn id=692 lang=golang
 *
 * [692] 前K个高频单词
 */

// @lc code=start
func Less(x, y string, dict map[string]int) bool {
	if dict[x] == dict[y] {
		return x < y
	}
	return dict[x] < dict[y]
}

func heapsort(str []string, root, end int, dict map[string]int) {
	for {
		child := 2*root + 1
		if child > end {
			break
		}
		if child+1 <= end && Less(str[child+1], str[child], dict) {
			child = child + 1
		}
		if Less(str[child], str[root], dict) {
			str[child], str[root] = str[root], str[child]
			root = child
		} else {
			break
		}
	}
}

func topKFrequent(words []string, k int) []string {
	n := len(words)
	dict := map[string]int{}
	ans := []string{}

	for i := 0; i < n; i++ {
		if _, ok := dict[words[i]]; !ok {
			dict[words[i]] = 1
			ans = append(ans, words[i])
		} else {
			dict[words[i]]++
		}
	}
	var hsort func([]string)
	hsort = func(ws []string) {
		end := len(ws) - 1
		for i := end / 2; i >= 0; i-- {
			heapsort(ws, i, end, dict)
		}
		for i := end; i >= 0; i-- {
			if Less(ws[0], ws[i], dict) {
				ws[0], ws[i] = ws[i], ws[0]
				heapsort(ws, 0, i-1, dict)
			}
		}
	}
	hsort(ans)
	sort.Slice(ans, func(i, j int) bool {
		if dict[ans[i]] == dict[ans[j]] {
			return ans[i] < ans[j]
		}
		return dict[ans[i]] > dict[ans[j]]
	})
	return ans[:k]
}

// @lc code=end

