package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// IPv4 lists string
// sort

// Less
func Less(ip1, ip2 string) bool {
	var (
		err    error
		ip1Ele int
		ip2Ele int
	)

	ip1Slice := strings.Split(ip1, ".")
	ip2Slice := strings.Split(ip2, ".")

	for i := 0; i < len(ip1Slice); i++ {
		// Convert
		ip1Ele, err = strconv.Atoi(ip1Slice[i])
		if err != nil {
			log.Println(err)
			return false
		}
		ip2Ele, err = strconv.Atoi(ip2Slice[i])
		if err != nil {
			log.Println(err)
			return false
		}
		// Compare
		if ip1Ele == ip2Ele {
			continue
		}

		if ip1Ele < ip2Ele {
			return true
		} else {
			break
		}
	}
	return false
}

// QuickSort Entry
func QuickSort(ips []string) {
	qsort(ips, 0, len(ips)-1)
}

// QuickSort Function
func qsort(slice []string, left, right int) {
	var (
		start int
		end   int
	)
	if left >= right {
		return
	}
	standard := slice[right]
	start = left
	end = right - 1
	for start < end {
		for Less(slice[start], standard) && start < end {
			start = start + 1
		}

		for !Less(slice[end], standard) && start < end {
			end = end - 1
		}
		slice[start], slice[end] = slice[end], slice[start]
	}
	// slice[right] is not compared
	if !Less(slice[start], slice[right]) {
		slice[start], slice[right] = slice[right], slice[start]
	} else {
		start = start + 1
	}
	if start == 0 {
		qsort(slice, start+1, right)
	}
	qsort(slice, left, start-1)
}

func main() {
	var (
		// n   int
		// ip  string
		ips []string
	)
	ips = []string{"11.255.255.255", "110.255.255.10", "25.140.125.123", "192.178.21.21"}
	// fmt.Scanln("%d", &n)
	// for i := 0; i < n; i++ {
	// 	fmt.Scanln("%s", &ip)
	// 	ips = append(ips, ip)
	// }
	// fmt.Println(n, ips)
	QuickSort(ips)
	fmt.Println(ips)
}
