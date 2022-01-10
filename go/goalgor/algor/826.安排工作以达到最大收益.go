/*
 * @lc app=leetcode.cn id=826 lang=golang
 *
 * [826] 安排工作以达到最大收益
 */

// @lc code=start

type job struct {
	id   int
	hard int
	prof int
}

func Less(job1, job2 job) bool {
	if job1.hard == job2.hard {
		return job1.prof > job2.prof
	}
	return job1.hard < job2.hard
}

func heap(jobs []job, root, end int) {
	for {
		child := 2*root + 1
		if child > end {
			break
		}
		if child+1 <= end && Less(jobs[child], jobs[child+1]) {
			child = child + 1
		}
		if Less(jobs[root], jobs[child]) {
			jobs[root], jobs[child] = jobs[child], jobs[root]
			root = child
		} else {
			break
		}
	}
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	n := len(profit)
	m := len(worker)
	var (
		find                  bool
		left, right, mid, ans int
		jobs                  []job
	)
	jobs = []job{}
	for i := 0; i < n; i++ {
		jobs = append(jobs, job{
			id:   i,
			hard: difficulty[i],
			prof: profit[i],
		})
	}

	for i := (n - 1) / 2; i >= 0; i-- {
		heap(jobs, i, n-1)
	}
	for end := n - 1; end >= 0; end-- {
		if Less(jobs[end], jobs[0]) {
			jobs[0], jobs[end] = jobs[end], jobs[0]
			heap(jobs, 0, end-1)
		}
	}
	for i := 1; i < n; i++ {
		if jobs[i-1].prof > jobs[i].prof {
			jobs[i].prof = jobs[i-1].prof
		}
	}

	ans = 0
	for i := 0; i < m; i++ {
		if worker[i] >= jobs[n-1].hard {
			ans += jobs[n-1].prof
			continue
		} else if worker[i] < jobs[0].hard {
			continue
		}
		left, right = 0, n-1
		mid = 0
		find = false
		for left < right {
			mid = (left + right) / 2
			if jobs[mid].hard < worker[i] {
				left = mid + 1
			} else if jobs[mid].hard > worker[i] {
				right = mid
			} else {
				find = true
				break
			}
		}
		if !find {
			ans += jobs[left-1].prof
		} else {
			ans += jobs[mid].prof
		}
	}
	return ans
}

// @lc code=end

