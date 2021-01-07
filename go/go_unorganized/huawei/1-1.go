package main

import (
	"fmt"
	"strconv"
	"strings"
)

func processInput(input string) ([][]float64, []int, []float64, []string, error) {
	var bvals []float64
	var vvals []int
	var params [][]float64
	res := strings.Split(input, ";")
	// Get the >, <, =
	spliter := strings.Split(res[len(res)-1], ",")
	// Get the target value
	b := strings.Split(res[len(res)-2], ",")
	for idx := range b {
		bval, err := strconv.ParseFloat(b[idx], 64)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		bvals = append(bvals, bval)
	}
	// Get the variants
	v := strings.Split(res[len(res)-3], ",")
	for idx := range v {
		vval, err := strconv.Atoi(v[idx])
		if err != nil {
			return nil, nil, nil, nil, err
		}
		vvals = append(vvals, vval)
	}
	// Get the params
	for i := 0; i < len(res)-3; i++ {
		temp_params := []float64{}
		temp_resi := strings.Split(res[i], ",")
		for idx := range temp_resi {
			temp_resival, err := strconv.ParseFloat(temp_resi[idx], 64)
			if err != nil {
				return nil, nil, nil, nil, err
			}
			temp_params = append(temp_params, temp_resival)
		}
		params = append(params, temp_params)
	}
	return params, vvals, bvals, spliter, nil
}

func getMax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func checkSpliter(s string, left, right float64) bool {
	if strings.Compare(s, "<=") == 0 {
		if left <= right {
			return true
		}
		return false
	}
	if strings.Compare(s, "<") == 0 {
		if left < right {
			return true
		}
		return false
	}
	if strings.Compare(s, "==") == 0 {
		if left == right {
			return true
		}
		return false
	}
	if strings.Compare(s, ">") == 0 {
		if left > right {
			return true
		}
		return false
	}
	if strings.Compare(s, ">=") == 0 {
		if left >= right {
			return true
		}
	}
	return false
}

func checkExists(params [][]float64, v []int, b []float64, spliter []string) (bool, int) {
	var leftVal []float64
	var maxSub float64
	for idx := range params {
		var val float64
		for i := 0; i < len(params[idx]); i++ {
			val += params[idx][i] * float64(v[i])
		}
		leftVal = append(leftVal, val)
	}
	count := 0
	maxSub = leftVal[0] - b[0]
	for idx := range b {
		if checkSpliter(spliter[idx], leftVal[idx], b[idx]) {
			count++
		}
		maxSub = getMax(maxSub, leftVal[idx]-b[idx])
	}
	if count == len(b) {
		return true, int(maxSub)
	}
	return false, int(maxSub)
}

func main1() {
	var input string
	fmt.Scanln(&input)
	p, v, b, spliter, err := processInput(input)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	ok, sub := checkExists(p, v, b, spliter)
	fmt.Println(ok, sub)
	return
}
