// /n means one layer deep
// /t means the position of the layer

var (
	maxLen int
	curLen int
	stk    []string
)

func lengthLongestPath(input string) int {
	names := strings.Split(input, "\n")
	stk = []string{}
	maxLen = 0
	curLen = 0

	for {
		if len(names) == 0 {
			break
		}
		numOfTab := strings.Count(names[0], "\t")
		for {
			if len(stk) == numOfTab {
				break
			}
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, names[0])
		if strings.Index(names[0], ".") != -1 {
			curLen = len(strings.Replace(strings.Join(stk, "/"), "\t", "", -1))
			if maxLen < curLen {
				maxLen = curLen
			}
		}
		names = names[1:]
	}

	return maxLen
}
