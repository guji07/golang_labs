package main

import (
	"fmt"
	"os"
)

func main() {
	var radius float64
	var solver Task1Solver = &TSolver{}
	fmt.Printf("type circle radius:\n")
	fmt.Scanf("%f", &radius)
	if radius <= 0 {
		fmt.Printf("incorrect data\n")
		os.Exit(1)
	}
	length, area := solver.CircleLengthAndArea(radius)
	fmt.Printf("length: %f area: %f\n", length, area)
}
