package main

import (
	"fmt"
	"os"
)

func main() {
	var radius, height float64
	var solver Task1Solver = &TSolver{}
	fmt.Printf("Type radius and height:")
	_, _ = fmt.Scanf("%f %f", &radius, &height)
	if radius <= 0 || height <= 0 {
		fmt.Printf("incorrect data\n")
		os.Exit(1)
	}
	fmt.Printf("cylinder volume: %f\n", solver.CylinderVolumeByRadiusAndHeight(radius, height))
	fmt.Printf("cone volume: %f\n", solver.ConeVolumeByByRadiusAndHeight(radius, height))
}
