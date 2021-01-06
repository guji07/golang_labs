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

func CheckPrime(number int) bool {
	isPrime := true
	if number == 0 || number == 1 {
		isPrime = false
	} else {
		for i := 2; i <= number/2; i++ {
			if number%i == 0 {
				isPrime = false
				break
			}
		}
	}
	return isPrime
}

func findAllPrimesBelowN(number int) {
	for i := 2; i < number; i++ {
		if CheckPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Printf("\n")
}
