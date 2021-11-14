/*
 * @lc app=leetcode.cn id=1953 lang=golang
 *
 * [1953] 你可以工作的最大周数
 */

// @lc code=start
func numberOfWeeks(milestones []int) int64 {
	var (
		ans int64
	)
	n := len(milestones)
	if n == 1 {
		return int64(1)
	}
	sort.Slice(milestones, func(i, j int) bool { return milestones[i] > milestones[j] })
	biggest := int64(milestones[0])
	sum := int64(0)
	for i := 1; i < n; i++ {
		sum = sum + int64(milestones[i])
	}
	if biggest > sum {
		ans = int64(2)*sum + 1
	} else if biggest == sum {
		ans = int64(2) * biggest
	} else {
		ans = biggest + sum
	}
	return ans
}

// @lc code=end
// type job struct {
// 	weeks int
// }

// type jobs []job

// func (js jobs) Less(i, j int) bool {
// 	return js[i].weeks > js[j].weeks
// }

// func (js jobs) Swap(i, j int) {
// 	js[i], js[j] = js[j], js[i]
// }

// func (js jobs) Len() int {
// 	return len(js)
// }

// func (js *jobs) Push(x interface{}) {
// 	*js = append(*js, x.(job))
// }

// func (js *jobs) Pop() interface{} {
// 	item := (*js)[len(*js)-1]
// 	*js = (*js)[:len(*js)-1]
// 	return item
// }

// func numberOfWeeks(milestones []int) int64 {
// 	var (
// 		sum     int64 = 0
// 		js      jobs
// 		lastJob *job
// 	)
// 	n := len(milestones)
// 	js = jobs{}

// 	// Initialize
// 	for i := 0; i < n; i++ {
// 		heap.Push(&js, job{
// 			weeks: milestones[i],
// 		})
// 	}
// 	sum++
// 	js[0].weeks--
// 	if js[0].weeks == 0 {
// 		lastJob = nil
// 	} else {
// 		tempJob := js[0]
// 		lastJob = &tempJob
// 	}
// 	heap.Pop(&js)

// 	// Begin
// 	for js.Len() != 0 && js[0].weeks > 0 {
// 		tempJob := lastJob
// 		sum++
// 		js[0].weeks--
// 		if js[0].weeks == 0 {
// 			lastJob = nil
// 		} else {
// 			copyJob := js[0]
// 			lastJob = &copyJob
// 		}
// 		heap.Pop(&js)
// 		if tempJob != nil {
// 			heap.Push(&js, *tempJob)
// 		}
// 	}

// 	return sum
// }
