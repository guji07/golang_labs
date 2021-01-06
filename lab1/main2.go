package main

import (
	"fmt"
	"os"
)

func main() {
	var edge float64
	var solver Task1Solver = &TSolver{}
	fmt.Printf("Type Cube edge size:")
	fmt.Scanf("%f", &edge)
	if edge <= 0 {
		fmt.Printf("incorrect data\n")
		os.Exit(1)
	}
	fmt.Printf("cube volume: %f\n", solver.CubeVolumeByEdge(edge))
}
