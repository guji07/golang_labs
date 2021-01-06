package main

import "fmt"

func main() {
	var A, B, C, D float64
	var solver Task1Solver = &TSolver{}
	fmt.Printf("type A B C D:\n")
	fmt.Scanf("%f %f %f %f", &A, &B, &C, &D)
	var value = solver.HouseSurfaceArea(A, D, B, C)
	if value > 0 {
		fmt.Printf("House surface area: %f", value)
	} else {
		fmt.Printf("invalid data")
	}
}
