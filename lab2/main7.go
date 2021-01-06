package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	fmt.Printf("type text:\n")
	r := bufio.NewReader(os.Stdin)
	line, _ := r.ReadString('\n')
	fmt.Printf("type char to count it's appereance in text:\n")
	var char byte
	fmt.Scanf("%c", &char)
	var counter int
	var counter2 int
	for _, value := range line {
		if unicode.IsLetter(value) {
			counter++
			if value == int32(char) {
				counter2++
			}
		}
	}
	fmt.Printf("all chars num: %d, '%c' num: %d\n", counter, char, counter2)
}
