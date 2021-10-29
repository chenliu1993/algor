package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/sys/unix"
)

var (
	execFifoFilename = "exec.fifo"
)

func createExecFifo(path string) error {
	fifoName := filepath.Join(path, execFifoFilename)
	if _, err := os.Stat(fifoName); err == nil {
		return fmt.Errorf("exec fifo %s already exists", fifoName)
	}
	oldMask := unix.Umask(0o000)
	if err := unix.Mkfifo(fifoName, 0o622); err != nil {
		unix.Umask(oldMask)
		return err
	}
	unix.Umask(oldMask)
	return nil
}

func readFromExecFifo(fifoPath string, flags int, block bool) error {
	if !block {
		flags |= unix.O_NONBLOCK
	}
	f, err := os.OpenFile(fifoPath, flags, 0)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	if len(data) <= 0 {
		return errors.New("cannot start an already running container")
	}
	fmt.Println(string(data))
	return nil
}

func writeToExecFifo(fifoPath string) error {
	fd, err := unix.Open(fifoPath, unix.O_WRONLY|unix.O_CLOEXEC, 0)
	if err != nil {
		return &os.PathError{Op: "open exec fifo", Path: fifoPath, Err: err}
	}
	if _, err := unix.Write(fd, []byte("here")); err != nil {
		return &os.PathError{Op: "write exec fifo", Path: fifoPath, Err: err}
	}
	_ = unix.Close(fd)
	// return os.Remove(fifoPath)
	return nil
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic to %v", r)
			return
		}
	}()
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fifoPath := filepath.Join(pwd, execFifoFilename)
	if err := createExecFifo(pwd); err != nil {
		panic(err)
	}

	go func(fifoPath string) {
		time.Sleep(time.Second)
		if err := writeToExecFifo(fifoPath); err != nil {
			panic(err)
		}

	}(fifoPath)

	if err := readFromExecFifo(fifoPath, os.O_RDONLY, true); err != nil {
		panic(err)
	}
}
