package main

import (
	"fmt"
	"strings"
)

func isInterleave(s1 string, s2 string, s3 string) bool {
	n1 := len(s1)
	n2 := len(s2)
	if n1+n2 != len(s3) {
		return false
	}
	record := make([][]bool, n1+1)
	for i := 0; i < n1+1; i++ {
		record[i] = make([]bool, n2+1)
		if strings.Compare(s1[:i], s3[0:i]) == 0 {
			record[i][0] = true
		} else {
			record[i][0] = false
		}
	}
	for i := 0; i < n2+1; i++ {
		if strings.Compare(s2[:i], s3[:i]) == 0 {
			record[0][i] = true
		} else {
			record[0][i] = false
		}
	}

	for i := 1; i < n1+1; i++ {
		for j := 1; j < n2+1; j++ {
			if s3[i+j-1] == s1[i-1] && s3[i+j-1] != s2[j-1] {
				record[i][j] = record[i-1][j]
			} else if s3[i+j-1] == s2[j-1] && s3[i+j-1] != s1[i-1] {
				record[i][j] = record[i][j-1]
			} else if s3[i+j-1] == s1[i-1] && s3[i+j-1] == s2[j-1] {
				record[i][j] = record[i-1][j] || record[i][j-1]
			} else {
				record[i][j] = false
			}
		}
	}
	return record[n1][n2]
}
func main() {
	s1 := "aabc"
	s2 := "abad"
	s3 := "aabadabc"
	fmt.Println(isInterleave(s1, s2, s3))
}
