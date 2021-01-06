package main

import "fmt"

func main() {
	var n1, n2, sum float64
	fmt.Scanf("%f %f", &n1, &n2)
	sum = n1 + n2
	fmt.Printf("\n%f\n", sum)
	fmt.Printf("%d\n", int(sum))
}
