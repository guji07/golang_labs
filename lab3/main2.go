package main

import (
	"fmt"
	"os"
)

func main() {
	var x float64
	var n int
	fmt.Println("type x and N:")
	_, _ = fmt.Scanf("%f%d", &x, &n)
	if x <= -1 || x >= 1 {
		fmt.Print("|x| not in range [0, 1]\n")
		os.Exit(1)
	}
	sum := rowSum(x, n)
	fmt.Printf("%f\n", sum)
}
