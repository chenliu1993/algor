package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	url := "https://www.facebook.com/favicon.ico"

	resp, err := http.Get(url)
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
