package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan string, 2)
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- "input 1 from routine 1"
		ch <- "input 2 from routine 1"
		close(ch)
	}()
	wg.Wait()
	for i := 0; i < 3; i++ {
		s := <-ch
		fmt.Println(len(s))
	}
}
