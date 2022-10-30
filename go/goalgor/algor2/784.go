package main

import "fmt"

func iterate(input string) {
	if len(input) == 0 {
		ans = append(ans, string(target))
		return
	}

	lastByte := input[0]
	if item, ok := dict[lastByte]; ok {
		old_target := make([]byte, len(target))
		copy(old_target, target)
		target = append(target, item)

		iterate(input[1:])
		target = make([]byte, len(old_target))
		copy(target, old_target)

		copy(old_target, target)
		target = append(target, lastByte)
		iterate(input[1:])
		target = make([]byte, len(old_target))
		copy(target, old_target)
	} else {
		target = append(target, input[0])
		iterate(input[1:])
	}
}

var ans []string
var target []byte
var dict map[byte]byte

func letterCasePermutation(s string) []string {
	dict = map[byte]byte{'a': 'A', 'A': 'a'}

	for i := 1; i < 26; i++ {
		dict['a'+byte(i)] = 'A' + byte(i)
		dict['A'+byte(i)] = 'a' + byte(i)
	}

	ans = []string{}
	target = []byte{}

	iterate(s)
	return ans

}

func main() {
	s := "3z4"
	fmt.Println(letterCasePermutation(s))
}
