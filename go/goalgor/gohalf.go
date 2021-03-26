package main

import (
	"fmt"

	algor "github.com/chenliu1993/algor/go/goalgor/algor"
)

func main() {
	src := []int{1, 5, 23, 7, 43, 7, 89, 2}
	fmt.Printf("%#v\n", src)
	algor.ShellSort(src)
	fmt.Println(src)
}
