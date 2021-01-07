package main

import "fmt"

func main() {
	var num, p int
	fmt.Printf("Type N:\n")
	_, _ = fmt.Scanf("%d", &num)
	for i := num; i != 0; i /= 10 {
		p = p*10 + i%10
	}
	if p == num {
		fmt.Printf("palindrome\n")
	} else {
		fmt.Printf("not palindrome\n")
	}
}
