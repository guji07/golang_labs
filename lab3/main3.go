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
	if x >= -1 && x <= 1 {
		fmt.Print("|x| not in range (-inf, -1) U (1, +inf)\n")
		os.Exit(1)
	}
	sum := secondRowSum(x, n)
	fmt.Printf("%.6f\n", sum)
}
