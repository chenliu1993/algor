package main

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	url, err := url.Parse("https://www.facebook.com/favicon.ico")
	if err != nil {
		panic(err)
	}
	request := http.Request{
		URL: url,
		Header: http.Header{
			"Content-Type": {"image/icon"},
		},
	}

	resp, err := http.DefaultClient.Do(&request)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("favicon.ico")
	if err != nil {
		panic(err)
	}
	file.Write(content)
}
