package main

import (
	"strconv"
	"strings"
)

func parseInts(s string) ([]int, error) {
	var (
		fields = strings.Fields(s)
		ints   = make([]int, len(fields))
		err    error
	)
	for i, f := range fields {
		ints[i], err = strconv.Atoi(f)
		if err != nil {
			return nil, err
		}
	}
	return ints, nil
}

func rowSum(x float64, n int) float64 {
	var sum float64 = 1
	if n <= 0 {
		return 0
	} else if n == 1 {
		return sum
	}
	for i := 2; i <= n; i++ {
		sum += x / float64(i)
		x *= x
	}
	return sum
}
