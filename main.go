package main

import (
	"fmt"
)

func main() {
	board := [][]byte{{"E", "E", "E", "E", "E"}}
	click := []int{3, 0}
	fmt.Println(updateBoard(board, click))
}
