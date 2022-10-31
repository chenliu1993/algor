package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}

	letterCh, numberCh := make(chan bool), make(chan bool)

	go func() {
		i := 0
		for {
			select {
			case <-numberCh:
				fmt.Println(i)
				i++
				letterCh <- true
			default:
				break
			}
		}
	}()

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		i := 0
		for {

			if i == 100 {
				wg.Done()
			}
			select {
			case <-letterCh:
				fmt.Println(string('a' + byte(i)))
				i++
				numberCh <- true
			default:
				break
			}
		}
	}(&wg)

	numberCh <- true
	wg.Wait()

}
