package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var counter int
	fmt.Printf("\ntype numbers:\n")
	r := bufio.NewReader(os.Stdin)
	line, _ := r.ReadString('\n')
	ints, _ := parseInts(line)
	for _, value := range ints {
		if value == 0 {
			counter++
		}
	}
	fmt.Printf("zero-values num: %d\n", counter)
}
