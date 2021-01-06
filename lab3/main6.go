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
		if value > max {
			max = value
		}
	}
	fmt.Printf("max: %d\n", max)
}
