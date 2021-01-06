package main

import "fmt"

func main() {
	var num1, num2, num3 float64
	var solver Task1Solver = &TSolver{}
	fmt.Printf("Type 3 numbers:\n")
	_, _ = fmt.Scanf("%f %f %f", &num1, &num2, &num3)
	fmt.Printf("\nmin: %.3f\n", solver.ThreeNumsMin(num1, num2, num3))
}
