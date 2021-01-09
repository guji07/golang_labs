package main

import (
	"fmt"
	"os"
)

func main() {
	var N, M int
	fmt.Printf("type N and M:\n")
	fmt.Scanf("%d %d", &N, &M)
	if N <= M {
		fmt.Println("wrong input: N <= M")
		os.Exit(1)
	}
	for i := N; i < M; i++ {
		if CheckPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Printf("\n")
}
