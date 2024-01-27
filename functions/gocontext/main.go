package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	ch := make(chan int, 1)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(ch chan int) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("context done")
				return
			case val := <-ch:
				fmt.Println("val:", val)
			}
		}
	}(ch)

	for i := 0; i < 111; i++ {
		ch <- i
	}

	wg.Wait()

}
