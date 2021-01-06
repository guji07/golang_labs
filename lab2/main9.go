package main

import "fmt"

func main() {
	var n, div = 2, 2
	fmt.Printf("type num:")
	fmt.Scanf("%d", &n)
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
