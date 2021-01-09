package main

import (
	"fmt"
	"os"
)

func main() {
	var n, div = 2, 2
	fmt.Printf("type num:\n")
	fmt.Scanf("%d", &n)
	if n == 0 {
		fmt.Println("num = 0")
		os.Exit(1)
	}
	fmt.Printf("num = 1")
	for ; n > 1; div++ {
		for ; n%div == 0; n /= div {
			{
				fmt.Printf(" * %d", div)
			}
		}
	}
	fmt.Printf("\n")
}
