package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	fmt.Printf("type numbers:\n")
	r := bufio.NewReader(os.Stdin)
	line, _ := r.ReadString('\n')
	ints, _ := parseInts(line)
	sort.Ints(ints)
	if len(ints) > 1 {
		fmt.Printf("%d %d\n", ints[0], ints[1])
	}
}
