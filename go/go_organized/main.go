package main

import (
	"fmt"

	"github.com/chenliu1993/algor/go/go_organized/utils"
)

func main() {
	ip1 := "192.168.1.1"
	ip2 := "192.168.2.1"
	mask := "255.255.255.0"
	ok, ip := utils.CompareNet(ip1, ip2, mask)
	fmt.Println(ok, ip)
	return
}
