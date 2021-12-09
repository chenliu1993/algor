package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
)

const (
	proxyURL = "http://{{ip}}"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		doSendRequests()
		wg.Done()
	}()

	wg.Wait()
}

func doSendRequests() {
	// client := newHTTPClient(nil)
	client := newHTTPClient(proxy)

	round := 0
	for {
		round += 1
		for i := 0; i < 5; i++ {
			go func(r int) {
				start := time.Now()
				var err error
				defer func() {
					duration := time.Since(start)
					fmt.Printf("Round: %d, A: %v, duration: %v\n", r, err, duration)
				}()
				req := newRequest("https://www.google.com")
				resp, err := client.Do(req)
				if err == nil {
					io.Copy(ioutil.Discard, resp.Body)
					resp.Body.Close()
				} else {
					fmt.Printf("A: %v\n", err)
				}
			}(round)
		}

		time.Sleep(100 * time.Second)
	}
}

type TLSDialer struct {
	net.Dialer
}

var (
	config = &tls.Config{InsecureSkipVerify: true}
)

func (dialer *TLSDialer) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	return tls.DialWithDialer(&dialer.Dialer, network, addr, config)
}

var (
	DefaultDialer = net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}

	DefaultTLSDialer = TLSDialer{DefaultDialer}
)

func proxy(req *http.Request) (*url.URL, error) {
	uri, err := url.Parse(proxyURL)
	if err != nil {
		return nil, nil
	}
	return uri, nil
}

func newHTTPClient(proxy func(*http.Request) (*url.URL, error)) *http.Client {
	t := &http.Transport{
		DialContext:           (&DefaultDialer).DialContext,
		DisableKeepAlives:     false,
		DisableCompression:    false,
		ForceAttemptHTTP2:     false,
		MaxIdleConns:          5000,
		MaxIdleConnsPerHost:   1000,
		IdleConnTimeout:       0 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	if proxy == nil {
		t.DialTLSContext = (&DefaultTLSDialer).DialContext
	} else {
		t.Proxy = proxy
		t.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	return &http.Client{Transport: t}
}

func newRequest(uri string) *http.Request {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	req, _ := http.NewRequestWithContext(ctx, "GET", uri, nil)

	return req
}
