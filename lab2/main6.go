package main

import (
	"fmt"
)

func main() {
	var alphabet = "abcdefghijklmnopqrstuvwxyz"
	fmt.Printf("%s\n", alphabet)
	for value, _ := range alphabet[1:] {
		fmt.Printf("%c\n", value+'a')
	}
}
