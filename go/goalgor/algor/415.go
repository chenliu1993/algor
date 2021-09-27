func add(num1 string, num2 string) string {
	n1 := len(num1)
	n2 := len(num2)
	carry := 0
	sum := 0
	i := n1 - 1
	j := n2 - 1
	var ans string
	for i >= 0 && j >= 0 {
		num1Int, err := strconv.Atoi(string(num1[i]))
		if err != nil {
			panic(err)
		}

		num2Int, err := strconv.Atoi(string(num2[j]))
		if err != nil {
			panic(err)
		}
		sum = num1Int + num2Int + carry
		sum, carry = sum%10, sum/10
		ans = strconv.Itoa(sum) + ans
		i = i - 1
		j = j - 1
	}
	for i >= 0 {
		sum = 0
		num1Int, err := strconv.Atoi(string(num1[i]))
		if err != nil {
			panic(err)
		}
		sum = num1Int + carry
		sum, carry = sum%10, sum/10
		ans = strconv.Itoa(sum) + ans
		i = i - 1
	}
	for j >= 0 {
		sum = 0
		num2Int, err := strconv.Atoi(string(num2[j]))
		if err != nil {
			panic(err)
		}
		sum = num2Int + carry
		sum, carry = sum%10, sum/10
		ans = strconv.Itoa(sum) + ans
		j = j - 1
	}
	if carry != 0 {
		ans = strconv.Itoa(carry) + ans
	}
	return ans
}