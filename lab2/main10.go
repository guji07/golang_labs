package main

import "fmt"

func main() {
	var N, M int
	fmt.Printf("type N and M:\n")
	fmt.Scanf("%d %d", &N, &M)
	for i := N; i < M; i++ {
		if CheckPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Printf("\n")
}
