package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Pair struct {
	value string
	time  int
}

func isNumber(s string) bool {
	is, err := regexp.MatchString("^[a-zA-Z]$", s)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return !is
}

func addPair(ps []Pair, p Pair) []Pair {
	var i int
	if len(ps) < 1 {
		ps = append(ps, p)
		return ps
	}
	for i = 0; i < len(ps); i++ {
		fmt.Println(ps)
		if ps[i].time > p.time {
			newps := append([]Pair{}, ps[i:]...)
			ps = append(ps[:i], p)
			ps = append(ps, newps...)
			break
		} else if ps[i].time < p.time {
			continue
		} else {
			if strings.Compare(ps[i].value, p.value) >= 0 {
				newps := append([]Pair{}, ps[i:]...)
				ps = append(ps[:i], p)
				ps = append(ps, newps...)
				break
			} else if strings.Compare(ps[i].value, p.value) < 0 {
				continue
			}
		}
	}
	if i == len(ps) {
		ps = append(ps, p)
		return ps
	}
	return ps
}

func splitIntoPairs(input string) []Pair {
	var (
		former int
		later  int
		pairs  []Pair
	)
	length := len(input)
	former = 0
	for later = former + 1; later < length; later++ {
		if isNumber(input[later : later+1]) {
			time, err := strconv.Atoi(input[later : later+1])
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			p := Pair{
				value: input[former:later],
				time:  time,
			}
			former = later + 1
			// pairs = append(pairs, p)
			pairs = addPair(pairs, p)
		}
	}
	return pairs
}

func Deploy(p []Pair) string {
	var rcs []string
	length := len(p)
	for i := 0; i < length; i++ {
		rc := ""
		for j := 0; j < p[i].time; j++ {
			rc = rc + p[i].value
		}
		rcs = append(rcs, rc)
	}
	return strings.Join(rcs, "")
}

func main() {
	var input string
	// fmt.Scanf("%s", &input)
	input = "a3b2c1"
	fmt.Println(Deploy(splitIntoPairs(input)))
	return
}
