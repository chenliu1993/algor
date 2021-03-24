package algor

// isSameSlice just does a simple comparsion.
func isSameSlice(src, dst []int) bool {
	if len(src) != len(dst) {
		return false
	}
	for i := range src {
		if src[i] != dst[i] {
			return false
		}
	}
	return true
}

// InsertSort: Input must be a slice or map otherwise copy-use will not affect the original param.
// TODO Why not working
func InsertSort(src []int) {
	var (
		i, j int
	)

	n := len(src)
	if n == 1 {
		return
	}

	for i = 1; i < n; i++ {
		for j = i; j > 0 && src[j-1] > src[j]; j-- {
			src[j], src[j-1] = src[j-1], src[j]
		}
	}
}
