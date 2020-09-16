package utils

import (
	"fmt"
	"io"
	"net"
	_ "net/http/pprof"
	"os"
	"sync"
	"sync/atomic"
	"syscall"
)

const (
	MaxBufferLen = 100
)

var BufPool = sync.Pool{
	New: func() interface{} {
		// TODO support unlinited buffer
		buffer := make([]byte, 10*1024*1024)
		return &buffer
	},
}

type udpReader struct {
	conn   *net.UDPConn
	closed uintptr
	Peer   *udpWriter
	Count  int64
}

func (u *udpReader) Read(p []byte) (int, error) {
	u.Count++
	n, err := u.conn.Read(p)
	if err == io.EOF {
		if atomic.LoadUintptr(&u.closed) != 0 {
			err = u.Peer.realClose()
			if err != nil {
				fmt.Println(err)
			}
			return 0, io.EOF
		}
	}
	return n, err
}

func (u *udpReader) Close() error {
	return u.conn.Close()
}

type udpWriter struct {
	conn   *net.UDPConn
	Peer   *udpReader
	Counts chan int64
	Count  int64
	// quit   chan bool
}

func (u *udpWriter) Write(b []byte) (int, error) {
	var err error
	fmt.Println("here", len(b))
	for i := 0; ; i++ {
		if len(b) < MaxBufferLen*(i+1) {
			if len(b) == MaxBufferLen*i {
				break
			}
			_, err = u.conn.Write(b[MaxBufferLen*i:])
			u.Count++
			u.Counts <- u.Count
			u.Counts <- -1
			break
		}
		_, err = u.conn.Write(b[MaxBufferLen*i : MaxBufferLen*(i+1)])
		if err != nil {
			// Something goes wrong here which may lead to write inconsistency, stop write.
			break
		}
		u.Count++
		u.Counts <- u.Count
	}
	if err != nil {
		errNet, ok := err.(*net.OpError)
		if ok {
			errSysCall, ok := errNet.Err.(*os.SyscallError)
			if ok {
				errSys, ok := errSysCall.Err.(syscall.Errno)
				if ok {
					if errSys == syscall.ECONNREFUSED {
						if err = u.Close(); err != nil {
							return len(b), err
						}
						return len(b), nil
					}
				}
			}
		}
	}
	return len(b), err
}

func (u *udpWriter) Close() error {
	atomic.StoreUintptr(&u.Peer.closed, 1)
	_, err := u.conn.Write([]byte{})
	if err != nil {
		fmt.Println(err)
	}
	return u.conn.Close()
}

func (u *udpWriter) realClose() error {
	_, err := u.conn.Write([]byte{})
	if err != nil {
		fmt.Println(err)
	}
	return u.conn.Close()
}
func UdpjobPipe(port int) (*udpReader, *udpWriter, error) {
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("127.0.0.1"),
	}
	read, err := net.ListenUDP("udp", &addr)
	if err != nil {
		return nil, nil, err
	}
	if err := read.SetReadBuffer(1024 * 1024); err != nil {
		return nil, nil, err
	}
	other := read.LocalAddr()
	local, ok := other.(*net.UDPAddr)
	if !ok {
		read.Close() // nolint: gosec
		return nil, nil, fmt.Errorf("failed to retrieve UDPAddr from %v (%T)", addr, addr)
	}

	write, err := net.DialUDP("udp", nil, local)
	if err != nil {
		read.Close() // nolint: gosec
		return nil, nil, err
	}
	reader := &udpReader{
		conn:  read,
		Count: int64(0),
	}
	ch := make(chan int64)
	// quitCh := make(chan bool)
	writer := &udpWriter{
		conn:   write,
		Peer:   reader,
		Count:  int64(0),
		Counts: ch,
		// quit:   quitCh,
	}
	reader.Peer = writer
	return reader, writer, nil
}

// Read for simulate read function.
func Read(r *udpReader) io.ReadCloser {
LOOP:
	for {
		select {
		case c := <-r.Peer.Counts:
			if c == -1 {
				fmt.Println("here1")
				return nil
			}
			if r.Count < c {
				break LOOP
			} else {
				fmt.Println("here2")
				return nil
			}
		}
	}
	return r
}
