package main

import (
	"fmt"
	"strings"
)

func removeComments(source []string) []string {
	sum := ""
	stk := ""
	ans := []string{}
	for i := 0; i < len(source); i++ {
		sum = sum + "##" + source[i]
	}
	sum = strings.TrimPrefix(sum, "##")
	sum = sum + "##"
	fmt.Println(sum)
	for i := 0; i < len(sum); i++ {
		if sum[i] == '/' {
			if sum[i+1] == '*' {
				i = strings.Index(sum[i+2:], "*/") + 3 + i
			} else if sum[i+1] == '/' {
				i = strings.Index(sum[i+2:], "##") + 1 + i
			} else {
				stk = stk + string(sum[i])
			}
		} else {
			stk = stk + string(sum[i])
		}
	}
	fmt.Println(stk)
	for strings.Contains(stk, "##") {
		idx := strings.Index(stk, "##")
		if idx == 0 {
			stk = stk[idx+2:]
			continue
		}
		ans = append(ans, stk[:idx])
		stk = stk[idx+2:]
	}
	if len(stk) != 0 {
		ans = append(ans, stk)
	}
	return ans
}
func main() {
	// source := []string{"a/*comment", "line", "more_comment*/b"}
	// source := []string{"/*Test program */", "int main()", "{ ", "  // variable declaration ", "int a, b, c;", "/* This is a test", "   multiline  ", "   comment for ", "   testing */", "a = b + c;", "}"}
	// source := []string{"void func(int k) {", "// this function does nothing /*", "   k = k*2/4;", "   k = k/2;*/", "}"}
	source := []string{"a//*b//*c", "blank", "d//*e/*/f"}
	// expected := []string{"int main()", "{ ", "  ", "int a, b, c;", "a = b + c;", "}"}
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	ans := removeComments(source)
	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i], i)
	}
	// fmt.Println(reflect.DeepEqual(removeComments(source), expected))
	// fmt.Println(removeComments(source))
}
