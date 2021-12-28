package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to file")

func init() {
	flag.Parse()
}

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

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC()                                      // GC，获取最新的数据信息
		if err := pprof.WriteHeapProfile(f); err != nil { // 写入内存信息
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go calcSum(&wg, i)
	}
	wg.Wait()
}
