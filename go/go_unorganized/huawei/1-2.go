package main

import (
	"fmt"
	"strings"
)

func findPosi(father, child string) {
	posi := strings.Index(father, child)
	if posi == -1 {
		fmt.Println("No")
		return
	}
	fmt.Println(posi + 1)
}

func main1() {
	var father, child string
	// father = "AVERDXIVYERDIAN"
	// child = "RDXI"
	fmt.Scanln(&father)
	fmt.Scanln(&child)
	findPosi(father, child)
	return
}
