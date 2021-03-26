package algor

func ShellSort(src []int) {
	var increment int
	n := len(src)
	if n < 2 {
		return
	}

	for increment = n / 2; increment > 0; increment = increment / 2 {
		for i := increment; i < n; i++ {
			for j := i; j >= increment; j -= increment {
				if src[j] < src[j-increment] {
					src[j], src[j-increment] = src[j-increment], src[j]
				}
			}
		}
	}
}
