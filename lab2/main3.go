package main

import "fmt"

func main() {
	var num1, num2, num3 float64
	var solver = &T2Solver{}
	fmt.Printf("type 3 numbers, where 1 is supposed to be max:\n")
	fmt.Scanf("%f %f %f", &num1, &num2, &num3)
	if solver.ThreeNumsMax(num1, num2, num3) {
		fmt.Printf("first num is really maximum number")
	} else {
		fmt.Printf("first number is not maximum")
	}
	if solver.ThreeNumsPositive(num1, num2, num3) {
		fmt.Printf("all nums are positive")
	} else {
		fmt.Printf("not all nums are positive")
	}
}
