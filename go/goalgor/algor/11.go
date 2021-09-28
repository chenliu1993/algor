// func getArea(line, hx,hy int) int {
//     if hx < hy {
//         return line * hx
//     }
//     return line * hy
// }

// func maxArea(height []int) int {
//     n := len(height)
//     maxArea := 0
//     record := make([][]int, n)
//     for i:=0;i<n;i++ {
//         record[i]=make([]int, n)
//         record[i][i]=0
//     }
//     for i:=1;i<n;i++ {
//         for j:=i-1;j>=0;j-- {
//             area := getArea(i-j, height[i], height[j])
//             if record[j+1][i] < area {
//                 record[j][i]= area
//             }else{
//                 record[j][i]=record[j+1][i]
//             }
//         }
//         if maxArea < record[0][i] {
//             maxArea = record[0][i]
//         }
//     }
//     return maxArea
// }

// func getArea(line, hx, hy int) int {
// 	if hx < hy {
// 		return line * hx
// 	}
// 	return line * hy
// }

// func maxArea(height []int) int {
//     n := len(height)
//     maxArea := 0
//     record := make([]int, n)
//     for i:=1;i<n;i++ {
//         lastMaxArea := 0
//         for j:=i-1;j>=0;j-- {
//             area := getArea(i-j, height[i], height[j])
//             if lastMaxArea < area {
//                 record[i]= area
//                 lastMaxArea=area
//             }else{
//                 record[i]=lastMaxArea
//             }
//         }
//         if maxArea < record[i] {
//             maxArea = record[i]
//         }
//     }
//     return maxArea
// }

// Above over time limit

func maxArea(height []int) int {
	n := len(height)
	maxArea := 0
	left := 0
	right := n - 1
	for left < right {
		if height[left] < height[right] {
			if maxArea < height[left]*(right-left) {
				maxArea = height[left] * (right - left)
			}
			left = left + 1
			continue
		}
		if maxArea < height[right]*(right-left) {
			maxArea = height[right] * (right - left)
		}
		right = right - 1
	}
	return maxArea
}