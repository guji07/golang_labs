package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 100; i < 1000; i++ {
		if IsPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
}

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
