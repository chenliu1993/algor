/*
 * @lc app=leetcode.cn id=649 lang=golang
 *
 * [649] Dota2 参议院
 */

// @lc code=start
var dict = map[string]string{
	"R": "Radiant",
	"D": "Dire",
}

func predictPartyVictory(senate string) string {
	n := len(senate)
	if n == 1 {
		return dict[senate]
	}
	var (
		id         string
		prohibited []int
	)
	for {
		n = len(senate)
		prohibited = make([]int, len(senate))
		for i := 0; i < n; i++ {
			if prohibited[i] == 1 {
				continue
			}
			j := i + 1
			j = j % n
			for j != i {
				if prohibited[j] != 1 {
					if senate[j] != senate[i] {
						prohibited[j] = 1
						break
					}
				}
				j = j + 1
				j = j % n
			}
		}
		tempSenates, rcount, dcount := "", 0, 0
		for i := 0; i < n; i++ {
			if prohibited[i] != 1 {
				if senate[i] == 'R' {
					rcount++
				} else {
					dcount++
				}
				tempSenates += string(senate[i])
			}
		}
		senate = tempSenates
		if rcount == n {
			id = "R"
			break
		} else if dcount == n {
			id = "D"
			break
		}
	}
	return dict[id]
}

// @lc code=end

// func predictPartyVictory(senate string) string {
// 	n := len(senate)
// 	var (
// 		r, d []int
// 	)
// 	r = []int{}
// 	d = []int{}
// 	for i := 0; i < n; i++ {
// 		if senate[i] == 'R' {
// 			r = append(r, i)
// 		} else {
// 			d = append(d, i)
// 		}
// 	}
// 	for len(r) > 0 && len(d) > 0 {
// 		if r[0] < d[0] {
// 			r = append(r, r[0]+n)
// 		} else {
// 			d = append(d, d[0]+n)
// 		}
// 		r = r[1:]
// 		d = d[1:]
// 	}
// 	if len(r) > 0 {
// 		return "Radiant"
// 	}
// 	return "Dire"
// }