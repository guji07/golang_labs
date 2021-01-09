package main

import "fmt"

func main() {
	var n1, n2, sum float64
	fmt.Println("Type 2 numbers")
	fmt.Scanf("%f %f", &n1, &n2)
	sum = n1 + n2
	fmt.Printf("sum: %f\n", sum)
	fmt.Printf("integer part of sum: %d\n", int(sum))
}
