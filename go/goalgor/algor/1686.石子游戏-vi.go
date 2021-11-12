/*
 * @lc app=leetcode.cn id=1686 lang=golang
 *
 * [1686] 石子游戏 VI
 */

// @lc code=start
type pair struct {
	val   int
	index int
}

type pairs []pair

func (ps pairs) Less(i, j int) bool {
	if ps[i].val < ps[j].val {
		return false
	}
	return true
}

func (ps pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps pairs) Len() int {
	return len(ps)
}

func stoneGameVI(aliceValues []int, bobValues []int) int {
	var (
		sum    int
		record pairs
	)

	for i := 0; i < len(aliceValues); i++ {
		record = append(record, pair{
			val:   aliceValues[i] + bobValues[i],
			index: i,
		})
	}
	sort.Sort(record)
	fmt.Println(record)
	sum = 0
	for i := 0; i < len(record); i++ {
		if i%2 == 0 {
			sum = sum + aliceValues[record[i].index]
		} else {
			sum = sum - bobValues[record[i].index]
		}
	}
	if sum > 0 {
		return 1
	} else if sum < 0 {
		return -1
	}
	return 0
}

// @lc code=end

