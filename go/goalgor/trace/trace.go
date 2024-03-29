package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func calcSum(w *sync.WaitGroup, idx int) {
	defer w.Done()
	var sum, n int64
	for ; n < 1000000000; n++ {
		sum += n
	}
	fmt.Println(idx, sum)
}

func main() {
	runtime.GOMAXPROCS(1)

	f, _ := os.OpenFile("./trace.out", os.O_RDWR|os.O_CREATE, 0755)
	defer f.Close()

	_ = trace.Start(f)
	defer trace.Stop()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go calcSum(&wg, i)
	}
	wg.Wait()
}
