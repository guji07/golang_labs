package main

import "fmt"

func main() {
	var N, M int
	fmt.Printf("Type N and M:\n")
	_, _ = fmt.Scanf("%d %d", &N, &M)
	if gcd(N, M) == 1 {
		fmt.Printf("true\n")
	} else {
		fmt.Printf("false\n")
	}
}

func gcd(a int, b int) int {
	if !(b > 0) {
		return a
	}
	return gcd(b, a%b)
}
