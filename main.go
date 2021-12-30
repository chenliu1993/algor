package main

import (
	"fmt"
	"github.com/chenliu1993/algor/pkg"
)

func main() {
	nums := []int{2, 34, 5, 1, 4, 45, 6}
	pkg.Sort(nums, 5)
	fmt.Println(nums)
}
