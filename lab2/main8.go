package main

import "fmt"

func main() {
	var num int
	fmt.Printf("Type N:\n")
	fmt.Scanf("%d", &num)
	if CheckPrime(num) {
		fmt.Printf("%d is a prime number\n", num)
	} else {
		fmt.Printf("%d is not a prime number\n", num)
	}
	fmt.Printf("all primes below %d: ", num)
	findAllPrimesBelowN(num)
}
