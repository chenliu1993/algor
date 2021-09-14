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