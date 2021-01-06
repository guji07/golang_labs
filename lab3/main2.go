package main

import "fmt"

func main() {
	var x float64
	var n int
	_, _ = fmt.Scanf("%f%d", &x, &n)
	if x <= -1 || x >= 1 {
		fmt.Print("|x| not in range [0, 1)")
	}
	sum := rowSum(x, n)
	fmt.Printf("%f", sum)
}
