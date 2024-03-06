/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	slen, plen := len(s), len(p)
	if slen < plen {
		return []int{}
	}

	ans := []int{}

	var sdict, pdict [26]int
	for i, ch := range p {
		sdict[s[i]-'a']++
		pdict[ch-'a']++
	}

	if sdict == pdict {
		ans = append(ans, 0)
	}

	for i, ch := range s[:slen-plen] {
		sdict[ch-'a']--
		// record the right edge
		sdict[s[i+plen]-'a']++
		if sdict == pdict {
			ans = append(ans, i+1)
		}

	}

	return ans

}

// @lc code=end

