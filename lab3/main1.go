package main

import "fmt"

func main() {
	var x float64
	var n int
	fmt.Println("type x and N:")
	_, _ = fmt.Scanf("%f%d", &x, &n)
	sum := rowSum(x, n)
	fmt.Printf("%f\n", sum)
}
