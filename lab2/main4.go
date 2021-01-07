package main

import "fmt"

func main() {
	var L1, L2, L3, T bool
	T = true
	L1 = false
	L2 = T && L1
	L1 = (!L2) || (!L1)
	L3 = L1 && L2 && T
	fmt.Printf("L1: %v L2: %v L3: %v \n", L1, L2, L3)
}
