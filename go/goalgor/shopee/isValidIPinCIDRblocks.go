package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func checkValidIP(blocks string, ip string) bool {
	cidrs := strings.Split(blocks, "/")
	masksVal, err := strconv.Atoi(cidrs[1])
	if err != nil {
		panic(err)
	}
	cidrFields := strings.Split(cidrs[0], ".")
	ipFields := strings.Split(ip, ".")
	if len(ipFields) != 4 || len(cidrFields) != 4 {
		panic("Not a valid IP address")
	}
	var (
		first, second int32
	)
	first = int32(0)
	second = int32(0)
	for i := 0; i < 4; i++ {
		fVal, err := strconv.Atoi(cidrFields[i])
		if err != nil {
			panic(err)
		}
		sVal, err := strconv.Atoi(ipFields[i])
		if err != nil {
			panic(err)
		}
		first = first + int32(fVal<<(8*(3-i)))
		second = second + int32(sVal<<(8*(3-i)))
	}
	return (first >> (32 - masksVal)) == (second >> (32 - masksVal))
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	blocks := "192.168.1.0/24"
	ip := "192.168.1.1"
	fmt.Println(checkValidIP(blocks, ip))
}
