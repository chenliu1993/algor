package algor

var (
	mid   int
	left  int
	right int
)

func hfind(src []int, tgt, left, right int) bool {
	if left > right {
		return false
	} else if left == right {
		if tgt == src[left] {
			return true
		}
		return false
	}
	mid = (left + right) / 2
	if tgt == src[mid] {
		return true
	} else if tgt < src[mid] {
		return hfind(src, tgt, left, mid)
	} else {
		return hfind(src, tgt, mid+1, right)
	}
}

func halfFind(src []int, tgt int) bool {
	return hfind(src, tgt, 0, len(src)-1)
}
