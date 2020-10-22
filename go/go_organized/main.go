package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("5 sec passed")
		default:
			time.Sleep(10 * time.Second)
			fmt.Println("10 sec passed")
		}
	}
}
