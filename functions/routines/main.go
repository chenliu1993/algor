package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func startInOrder(start int, inCh, outCh chan int) {
	defer wg.Done()
	for i := start; i < 101; {
		previousVal, ok := <-inCh
		if !ok {
			close(inCh)
			break
		}
		fmt.Println(previousVal)
		outCh <- i
		i = i + 3
	}
}

// one routine to print numbers, another one to print alphabets
func startInTurnAlpha(charCh, numCh chan struct{}, count int) {
	i := 65
	for {
		<-charCh

		if i > 91 {
			close(numCh)
			return
		}

		for j := 0; j < count; j++ {
			fmt.Println("char routine", string(i))
			i++
		}
		numCh <- struct{}{}
	}

}

func startInTurnNum(charCh, numCh chan struct{}, count int) {
	i := 1
	for {
		_, ok := <-numCh
		if !ok {
			// close(numCh)
			return
		}
		for j := 0; j < count; j++ {
			fmt.Println("num routine", i)
			i++
		}
		charCh <- struct{}{}
	}

}

func main() {

	n := 10
	chs := make([]chan struct{}, n)
	for i := 0; i < n; i++ {
		chs[i] = make(chan struct{})
	}

	char := 65
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				_, ok := <-chs[i]
				if !ok {
					return
				}
				if char > 90 {
					for _, ch := range chs {
						close(ch)
					}
					return
				}

				fmt.Println(fmt.Sprintf("routine %d", i), string(char))
				char++
				chs[(i+1)%n] <- struct{}{}
			}
		}(i)

	}
	chs[0] <- struct{}{}
	wg.Wait()

	/* numCh, charCh := make(chan struct{}), make(chan struct{})
	go func() {
		charCh <- struct{}{}
	}()
	wg.Add(2)
	go func() {
		defer wg.Done()
		startInTurnAlpha(charCh, numCh, 2)
	}()
	go func() {
		defer wg.Done()
		startInTurnNum(charCh, numCh, 2)
	}()
	wg.Wait() */
	/* ch1, ch2, ch3 := make(chan int, 1), make(chan int, 1), make(chan int, 1)

	wg.Add(3)
	ch1 <- 1
	go startInOrder(1, ch1, ch2)
	go startInOrder(2, ch2, ch3)
	go startInOrder(3, ch3, ch1)
	wg.Wait() */

}
