package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/chenliu1993/algor/go/go_organized/utils"
)

func main() {
	// runtime.SetMutexProfileFraction(1)
	// runtime.SetBlockProfileRate(1)
	// go func() {
	// 	log.Println(http.ListenAndServe(":8088", nil))
	// }()

	r1, w1, err := os.Pipe()
	if err != nil {
		fmt.Println(err)
	}
	r2, w2, err := os.Pipe()
	if err != nil {
		fmt.Println(err)
	}
	reader1, writer1, err := utils.UdpjobPipe(0)
	if err != nil {
		fmt.Println(err)
	}
	_, writer2, err := utils.UdpjobPipe(0)
	if err != nil {
		fmt.Println(err)
	}
	cmd := exec.Command("./log_test")
	cmd.Stdout = w1
	cmd.Stderr = w2
	copyClose := func(r *os.File, w io.WriteCloser) {
		fmt.Println("copy begins")
		_, err := io.Copy(w, r)
		fmt.Println("copy ends")
		if err != nil {
			fmt.Println("copy error", err)
		}
		if err := w.Close(); err != nil {
			fmt.Println(err)
		}
		if err := r.Close(); err != nil {
			fmt.Println(err)
		}
	}
	go copyClose(r1, writer1)
	go copyClose(r2, writer2)
	buf := make([]byte, utils.MaxBufferLen)
	go cmd.Run()
	// var wg sync.WaitGroup
	// wg.Add(1)
	go func() {
		// defer wg.Done()
		for {
			tempReader := utils.Read(reader1)
			if tempReader == nil {
				fmt.Println("problem")
				return
			}
			n, err := tempReader.Read(buf)
			if err != nil {
				fmt.Println("read err: ", err)
				return
			}
			fmt.Println(string(buf[:n]))
		}
	}()
	// wg.Wait()
	time.Sleep(10 * time.Second)
	return
}
