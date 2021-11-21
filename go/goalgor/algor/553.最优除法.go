/*
 * @lc app=leetcode.cn id=553 lang=golang
 *
 * [553] 最优除法
 */

// @lc code=start
type result struct {
	res  float64
	expr string
}

func optimalDivision(nums []int) string {
	n := len(nums)
	if n == 1 {
		return strconv.Itoa(nums[0])
	}
	input := ""
	for i := 0; i < n; i++ {
		input = input + strconv.Itoa(nums[i]) + "/"
	}
	input = input[:len(input)-1]
	var findRes func(string) []result

	findRes = func(s string) []result {
		if !strings.ContainsAny(s, "/") {
			if o, err := strconv.ParseFloat(s, 64); err != nil {
				return []result{
					{
						res:  float64(0),
						expr: "",
					},
				}
			} else {
				return []result{
					{
						res:  o,
						expr: s,
					},
				}
			}
		}
		ans := []result{}
		for i := 0; i < len(s); i++ {
			if s[i] != '/' {
				continue
			}
			left := findRes(s[:i])
			right := findRes(s[i+1:])
			for _, lval := range left {
				var lexpr string
				if !strings.ContainsAny(lval.expr, "/") {
					lexpr = s[:i]
				} else {
					lexpr = "(" + s[:i] + ")"
				}
				for _, rval := range right {
					var rexpr string
					if !strings.ContainsAny(rval.expr, "/") {
						rexpr = s[i+1:]
					} else {
						rexpr = "(" + s[i+1:] + ")"
					}
					tempRes := result{
						res:  lval.res / rval.res,
						expr: lexpr + "/" + rexpr,
					}
					ans = append(ans, tempRes)
				}
			}
		}
		return ans
	}
	output := findRes(input)
	sort.Slice(output, func(i, j int) bool {
		return output[i].res > output[j].res
	})
	return output[0].expr
}

// @lc code=end

