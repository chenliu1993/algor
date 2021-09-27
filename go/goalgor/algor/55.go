func canJump(nums []int) bool {
	n := len(nums)
	rightmost := 0
	for i := 0; i < n; i++ {
		if i <= rightmost {
			if rightmost < i+nums[i] {
				rightmost = i + nums[i]
			}
			if rightmost >= n-1 {
				return true
			}
		}

	}
	return false
}

// Add dp
// func canJump(nums []int) bool {
//     n := len(nums)
//     can := make([]int, n)
//     can[0]=1
//     for i:=0;i<n;i++ {
//         for j:=i-1;j>=0;j-- {
//             if can[j]==1 && (i-j) <= nums[j] {
//                 can[i]=1
//                 break
//             }
//         }
//     }
//     return can[n-1] == 1
// }