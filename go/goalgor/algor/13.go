// 通常情况下，罗马数字中小的数字在大的数字的右边。若输入的字符串满足该情况，那么可以将每个字符视作一个单独的值，累加每个字符对应的数值即可。

// 例如 \texttt{XXVII}XXVII 可视作 \texttt{X}+\texttt{X}+\texttt{V}+\texttt{I}+\texttt{I}=10+10+5+1+1=27X+X+V+I+I=10+10+5+1+1=27。

// 若存在小的数字在大的数字的左边的情况，根据规则需要减去小的数字。对于这种情况，我们也可以将每个字符视作一个单独的值，若一个数字右侧的数字比它大，则将该数字的符号取反。

// 例如 \texttt{XIV}XIV 可视作 \texttt{X}-\texttt{I}+\texttt{V}=10-1+5=14X−I+V=10−1+5=14。

var convertMappings = map[string]int{
	"I":  1,
	"IV": 4,
	"IX": 9,
	"V":  5,
	"X":  10,
	"XL": 40,
	"XC": 90,
	"L":  50,
	"C":  100,
	"CD": 400,
	"CM": 900,
	"D":  500,
	"M":  1000,
}

func romanToInt(s string) int {
	sum := 0
	ele := strings.Split(s, "")
	n := len(ele)
	for i := 0; i < n; i++ {
		if ele[i] == "I" && i < n-1 {
			if ele[i+1] == "V" {
				sum += convertMappings["IV"]
				i = i + 1
				continue
			}
			if ele[i+1] == "X" {
				sum += convertMappings["IX"]
				i = i + 1
				continue
			}
			sum += convertMappings["I"]
			continue
		}
		if ele[i] == "X" && i < n-1 {
			if ele[i+1] == "L" {
				sum += convertMappings["XL"]
				i = i + 1
				continue
			}
			if ele[i+1] == "C" {
				sum += convertMappings["XC"]
				i = i + 1
				continue
			}
			sum += convertMappings["X"]
			continue
		}
		if ele[i] == "C" && i < n-1 {
			if ele[i+1] == "D" {
				sum += convertMappings["CD"]
				i = i + 1
				continue
			}
			if ele[i+1] == "M" {
				sum += convertMappings["CM"]
				i = i + 1
				continue
			}
			sum += convertMappings["C"]
			continue
		}
		sum += convertMappings[ele[i]]
	}
	return sum
}