package main

import "fmt"

func main() {
	var A, B, C int
	A = 7
	B = 9
	C = 5
	booleanValue := (A > B) && (B == A+2) || !(C != B)
	fmt.Printf("booleanValue: %v\n", booleanValue)
	C = 9
	booleanValue = (A > B) && (B == A+2) || !(C != B)
	fmt.Printf("new booleanValue: %v\n", booleanValue)
}
