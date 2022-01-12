package main

import (
	"fmt"
)

func patternMatching(pattern string, value string) bool {
	plen := len(pattern)
	vlen := len(value)

	var (
		find  bool
		count map[byte]int
	)
	count = map[byte]int{
		'a': 0,
		'b': 0,
	}
	for i := 0; i < plen; i++ {
		count[pattern[i]]++
	}
	if count['a'] < count['b'] {
		count['a'], count['b'] = count['b'], count['a']
		temp := ""
		for i := 0; i < plen; i++ {
			if pattern[i] == 'a' {
				temp += "b"
			} else {
				temp += "a"
			}
		}
		pattern = temp
	}
	if plen == 0 {
		return false
	} else if vlen == 0 {
		return count['b'] == 0
	}

	for i := 0; i <= (vlen / count['a']); i++ {
		rest := (vlen - count['a']*i)
		if (count['b'] == 0 && rest == 0) || (count['b'] != 0 && rest%count['b'] == 0) {
			var j int
			if count['b'] == 0 {
				j = 0
			} else {
				j = rest / count['b']
			}
			pos := 0
			cur, valueA, valueB := "", "", ""
			no := false
			for k := 0; k < plen; k++ {
				if pattern[k] == 'a' {
					cur = value[pos : pos+i]
					if len(valueA) == 0 {
						valueA = cur
					} else if valueA != cur {
						no = true
						break
					}
					pos += i
				} else {
					cur = value[pos : pos+j]
					if len(valueB) == 0 {
						valueB = cur
					} else if valueB != cur {
						no = true
						break
					}
					pos += j
				}
			}
			if valueA != valueB && !no {
				find = true
				break
			}
		}
	}
	return find
}

func main() {
	pattern := "abba"
	value := "dogcatcatdog"
	fmt.Println(patternMatching(pattern, value))
}
