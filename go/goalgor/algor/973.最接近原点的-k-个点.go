/*
 * @lc app=leetcode.cn id=973 lang=golang
 *
 * [973] 最接近原点的 K 个点
 */

// @lc code=start
func EuclidDistance(x1, y1 float64, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, float64(2)) + math.Pow(y2-y1, float64(2)))
}

type Point struct {
	x, y float64
}

func Less(p1, p2 Point) bool {
	return EuclidDistance(0, 0, p1.x, p1.y) < EuclidDistance(0, 0, p2.x, p2.y)
}

func heapSort(points []Point) {
	end := len(points) - 1
	hsort := func(l, r int) {
		for {
			c := 2*l + 1
			if c > r {
				break
			}
			if c+1 <= r && Less(points[c], points[c+1]) {
				c = c + 1
			}
			if Less(points[l], points[c]) {
				points[l], points[c] = points[c], points[l]
				l = c
			} else {
				break
			}
		}
	}
	for root := end / 2; root >= 0; root-- {
		hsort(root, end)
	}

	for last := end; last >= 0; last-- {
		if Less(points[last], points[0]) {
			points[last], points[0] = points[0], points[last]
			hsort(0, last-1)
		}
	}
}

func quickSort(points []Point) {
	var qsort func(int, int)
	qsort = func(start, end int) {
		if start >= end {
			return
		}
		left := start
		right := end - 1
		standard := points[end]
		for left < right {
			for left < right && Less(points[left], standard) {
				left++
			}
			for left < right && Less(standard, points[right]) {
				right--
			}
			points[left], points[right] = points[right], points[left]
		}
		if Less(points[end], points[left]) {
			points[end], points[left] = points[left], points[end]
		} else {
			left++
		}
		if left != 0 {
			qsort(start, left-1)
		}
		qsort(left+1, end)
	}
	qsort(0, len(points)-1)
}

func shellSort(points []Point) {
	n := len(points)
	if n < 2 {
		return
	}
	for increment := n / 2; increment > 0; increment = increment / 2 {
		for i := increment; i < n; i++ {
			for j := i; j >= increment; j = j - increment {
				if Less(points[i], points[i-increment]) {
					points[i], points[i-increment] = points[i-increment], points[i]
				}
			}
		}
	}
}

func bubbleSort(points []Point) {
	n := len(points)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if Less(points[j], points[j-1]) {
				points[j], points[j-1] = points[j-1], points[j]
			}
		}
	}
}

func Sort(ps []Point, method string) {
	if method == "Heap" {
		heapSort(ps)
	} else if method == "Shell" {
		shellSort(ps)
	} else if method == "Quick" {
		quickSort(ps)
	} else {
		bubbleSort(ps)
	}
}

func kClosest(points [][]int, k int) [][]int {
	Points := []Point{}
	n := len(points)
	for i := 0; i < n; i++ {
		Points = append(Points, Point{
			x: float64(points[i][0]),
			y: float64(points[i][1]),
		})
	}
	Sort(Points, "Heap")
	ans := [][]int{}
	for i := 0; i < k; i++ {
		ans = append(ans, []int{int(Points[i].x), int(Points[i].y)})
	}
	return ans
}

// @lc code=end

