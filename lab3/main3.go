package main

import "fmt"

func main() {
	var x float64
	var n int
	_, _ = fmt.Scanf("%f%d", &x, &n)
	if x >= -1 && x <= 1 {
		fmt.Print("x not in range (-1, 1)")
	}
	sum := secondRowSum(x, n)
	fmt.Printf("%.6f\n", sum)
}

func secondRowSum(x float64, n int) float64 {
	var sum float64 = 1
	if n <= 0 {
		return 0
	} else if n == 1 {
		return sum
	}
	for i := 2; i <= n; i++ {
		sum += (-x) / float64(i)
		x *= -x
	}
	return sum
}
