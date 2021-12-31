package main

import (
	"fmt"
)

func predictPartyVictory(senate string) string {
	n := len(senate)
	var (
		r, d []int
	)
	r = []int{}
	d = []int{}
	for i := 0; i < n; i++ {
		if senate[i] == 'R' {
			r = append(r, i)
		} else {
			d = append(d, i)
		}
	}
	for len(r) > 0 && len(d) > 0 {
		if r[0] < d[0] {
			r = append(r, r[0]+n)
		} else {
			d = append(d, d[0]+n)
		}
		r = r[1:]
		d = d[1:]
	}
	if len(r) > 0 {
		return "Radiant"
	}
	return "Dire"
}

func main() {
	senate := "DDRRR"
	fmt.Println(predictPartyVictory(senate))
}
