package main

import (
	"fmt"
	"math"
)

type Task2Solver interface {
	roundFloat(num float64) int
	volumesSum(startRadius float64, deltaRadius float64, spheraNum int) float64
}

type T2Solver struct{}

func (T *T2Solver) roundFloat(num float64) int {
	return int(num)
}

func (T *T2Solver) volumesSum(startRadius float64, deltaRadius float64, spheraNum int) float64 {
	var volume float64 = 0.0
	for i := 0; i < spheraNum; i++ {
		radius := startRadius + (deltaRadius * float64(i))
		sphereVolume := 4 * math.Pi * math.Pow(radius, 3.0)
		fmt.Printf("sphere volume: %f radius: %f\n", sphereVolume, radius)
		volume += sphereVolume
	}
	return volume
}

func (T *T2Solver) ThreeNumsMax(supposedMax float64, num2 float64, num3 float64) bool {
	if supposedMax >= num2 && supposedMax >= num3 {
		return true
	}
	return false
}

func (T *T2Solver) ThreeNumsPositive(num1 float64, num2 float64, num3 float64) bool {
	if num1 >= 0 && num2 >= 0 && num3 >= 0 {
		return true
	}
	return false
}
