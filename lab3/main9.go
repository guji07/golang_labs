package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var max int
	fmt.Printf("type numbers:\n")
	r := bufio.NewReader(os.Stdin)
	line, _ := r.ReadString('\n')
	ints, _ := parseInts(line)
	max = math.MinInt32
	for _, value := range ints {
		if value > max && value < 0 {
			max = value
		}
	}
	fmt.Printf("max negative num: %d\n", max)
}
