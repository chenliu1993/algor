package algor

import "math"

func findSplit(points [][2]int) (int, int) {
	var (
		xmax = points[0][0]
		xmin = points[0][0]
		ymax = points[0][1]
		ymin = points[0][1]
	)
	for _, point := range points {
		if xmax < point[0] {
			xmax = point[0]
			continue
		}
		if ymax < point[1] {
			ymax = point[1]
			continue
		}
		if xmin > point[0] {
			xmin = point[0]
			continue
		}
		if ymin > point[1] {
			ymin = point[1]
			continue
		}
	}
	return (xmax + xmin) / 2, (ymax + ymin) / 2
}

func twoPointDis(p1, p2 [2]int) float64 {
	return math.Sqrt(float64((p1[0] - p2[0]) ^ 2 + (p1[1] - p2[1]) ^ 2))
}

// Time O(N^2)???
func minDistance(points [][2]int) float64 {
	xmid, _ := findSplit(points)
	// first calculate all the points in the left of the xmid,
	// then the points in the right of the split,
	// finally the distance which one end is in the left and the other is in the right.
	leftPoints := [][2]int{}
	rightPoints := [][2]int{}
	var (
		middleMinDistance = math.MaxFloat64
		rightMinDistance  = math.MaxFloat64
		leftMinDistance   = math.MaxFloat64
	)
	for _, point := range points {
		if point[0] <= xmid {
			leftPoints = append(leftPoints, point)
			continue
		}
		rightPoints = append(rightPoints, point)
	}
	if len(rightPoints) > 1 {
		rightMinDistance = minDistance(rightPoints)
	}
	if len(leftPoints) > 1 {
		leftMinDistance = minDistance(leftPoints)
	}

	for _, leftPoint := range leftPoints {
		for _, rightPoint := range rightPoints {
			middleMinDistance = math.Min(middleMinDistance, twoPointDis(leftPoint, rightPoint))
		}
	}
	return math.Min(middleMinDistance, math.Min(leftMinDistance, rightMinDistance))
}
