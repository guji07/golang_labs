package main

import (
	"fmt"
	"os"
)

func main() {
	var num int
	fmt.Printf("Type N:\n")
	_, _ = fmt.Scanf("%d", &num)
	if num <= 0 {
		fmt.Println("wrong input: num <= 0")
		os.Exit(1)
	}
	fmt.Printf("prime below %d: %d\n", num, findOnePrimeBelowN(num))
	if isPerfectNumber(num) {
		fmt.Printf("%d is perfect number\n", num)
	} else {
		fmt.Printf("%d is not perfect number\n", num)
	}
}
