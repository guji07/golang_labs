package main

import (
	"fmt"
	"math"
)

func roundFloat(num float64) int {
	return int(num)
}

func volumesSum(startRadius float64, deltaRadius float64, k float64) float64 {
	var volume float64 = 0.0
	var radius float64 = startRadius
	for i := 0; radius < k; i++ {
		sphereVolume := 4 * math.Pi * math.Pow(radius, 3.0)
		fmt.Printf("sphere volume: %.3f radius: %.3f\n", sphereVolume, radius)
		volume += sphereVolume
		radius = startRadius + (deltaRadius * float64(i))
	}
	return volume
}

func ThreeNumsMax(supposedMax float64, num2 float64, num3 float64) bool {
	if supposedMax >= num2 && supposedMax >= num3 {
		return true
	}
	return false
}

func ThreeNumsPositive(num1 float64, num2 float64, num3 float64) bool {
	if num1 >= 0 && num2 >= 0 && num3 >= 0 {
		return true
	}
	return false
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

func findOnePrimeBelowN(number int) int {
	var prime int
	for i := 2; i < number; i++ {
		if CheckPrime(i) {
			prime = i
		}
	}
	return prime
}

func isPerfectNumber(n int) bool {
	var sum = 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum = sum + i
		}
	}
	if sum == n {
		return true
	} else {
		return false
	}
}
