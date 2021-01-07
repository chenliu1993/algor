package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// CompareNet is an example for HuaWei from Internet
func CompareNet(ip1, ip2, mask string) (bool, string) {
	ip1Segs := strings.Split(ip1, ".")
	ip2Segs := strings.Split(ip2, ".")
	maskSegs := strings.Split(mask, ".")
	var ip1Mask string
	count := 0
	for idx := range ip1Segs {
		ip1SegsInt, _ := strconv.Atoi(ip1Segs[idx])
		ip2SegsInt, _ := strconv.Atoi(ip2Segs[idx])
		maskSegsInt, _ := strconv.Atoi(maskSegs[idx])
		if (ip1SegsInt & maskSegsInt) == (ip2SegsInt & maskSegsInt) {
			count++
		}
		ip1Mask += strconv.Itoa(ip1SegsInt & maskSegsInt)
		fmt.Println(strconv.Itoa(ip1SegsInt & maskSegsInt))
		ip1Mask += "."
	}
	if count != len(ip2Segs) {
		return false, ip1Mask[:len(ip1Mask)-1]
	}
	return true, ip1Mask[:len(ip1Mask)-1]
}
