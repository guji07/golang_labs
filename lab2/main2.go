package main

import "fmt"

func main() {
	var r, deltaR float64
	var k int
	fmt.Printf("type first sphere radius, deltaR and k:\n")
	fmt.Scanf("%f %f %d", &r, &deltaR, &k)
	var solver = &T2Solver{}
	var volume = solver.volumesSum(r, deltaR, k)
	fmt.Printf("%f\n", volume)
}
