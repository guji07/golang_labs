package main

import "fmt"

func main() {
	var x, y, x1, x2, y1, y2 float64
	var solver Task1Solver = &TSolver{}
	fmt.Printf("type coordinates: x, y, x1, x2, y1, y2\n")
	_, _ = fmt.Scanf("%f %f %f %f %f %f", &x, &y, &x1, &x2, &y1, &y2)
	value, err := solver.IsPointBelongToRectangle(&Rectangle{
		x1: x1,
		x2: x2,
		y1: y1,
		y2: y2,
	}, Point{
		x: x,
		y: y,
	})
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		fmt.Printf("answer: %t\n", value)
	}
}
