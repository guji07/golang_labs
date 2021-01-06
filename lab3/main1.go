package main

import "fmt"

func main() {
	var x float64
	var n int
	_, _ = fmt.Scanf("%f%d", &x, &n)
	sum := rowSum(x, n)
	fmt.Printf("%f", sum)
}
