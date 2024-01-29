package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern, err := regexp.Compile("^+$")
	if err != nil {
		panic(err)
	}

	if pattern.MatchString("child") {
		fmt.Println("Matched")
	}
}
