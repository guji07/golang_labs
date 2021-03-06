package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var max int
	fmt.Printf("\ntype numbers:\n")
	r := bufio.NewReader(os.Stdin)
	line, _ := r.ReadString('\n')
	ints, _ := parseInts(line)
	for _, value := range ints {
		if value > max && (value&1) == 0 {
			max = value
		}
	}
	fmt.Printf("max odd num: %d\n", max)
}
