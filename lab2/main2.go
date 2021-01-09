package main

import (
	"fmt"
	"os"
)

func main() {
	var r, deltaR, k float64
	fmt.Printf("type first sphere radius, deltaR and k:\n")
	fmt.Scanf("%f %f %f", &r, &deltaR, &k)
	if r > k {
		fmt.Printf("wrong input data: k < r\n")
		os.Exit(1)
	}
	var volume = volumesSum(r, deltaR, k)
	fmt.Printf("%f\n", volume)
}
