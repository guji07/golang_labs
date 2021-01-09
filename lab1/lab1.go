package main

import (
	"fmt"
	"math"
)

type Task1Solver interface {
	CircleLengthAndArea(radius float64) (length float64, area float64)
	CubeVolumeByEdge(edge float64) float64
	CylinderVolumeByRadiusAndHeight(radius float64, height float64) float64
	ConeVolumeByByRadiusAndHeight(radius float64, height float64) float64
	IsPointBelongToRectangle(rectangle *Rectangle, point Point) (bool, error)
	HouseSurfaceArea(height float64, width float64, windowHeight float64, windowWidth float64) float64
	ThreeNumsSum(num1 float64, num2 float64, num3 float64) float64
	ThreeNumsMin(num1 float64, num2 float64, num3 float64) float64
}

type TSolver struct{}

type Rectangle struct {
	x1 float64
	x2 float64
	y1 float64
	y2 float64
}

type Point struct {
	x float64
	y float64
}

func (t *TSolver) CircleLengthAndArea(radius float64) (length float64, area float64) {
	if radius <= 0 {
		return -1, -1
	}
	return 2 * math.Pi * radius, math.Pi * math.Pow(radius, 2)
}

func (t *TSolver) CubeVolumeByEdge(edge float64) float64 {
	if edge <= 0 {
		return -1
	}
	return math.Pow(edge, 3)
}

func (t *TSolver) CylinderVolumeByRadiusAndHeight(radius float64, height float64) float64 {
	if radius <= 0 || height <= 0 {
		return -1
	}
	return math.Pi * math.Pow(radius, 2) * height
}

func (t *TSolver) ConeVolumeByByRadiusAndHeight(radius float64, height float64) float64 {
	if radius <= 0 || height <= 0 {
		return -1
	}
	return t.CylinderVolumeByRadiusAndHeight(radius, height) * (1.0 / 3.0)
}

func (t *TSolver) IsPointBelongToRectangle(rectangle *Rectangle, point Point) (bool, error) {
	if rectangle.x1 >= rectangle.x2 || rectangle.y1 >= rectangle.y2 {
		return false, fmt.Errorf("invalid Rectangle")
	}
	return point.x >= rectangle.x1 &&
		point.x <= rectangle.x2 &&
		point.y >= rectangle.y1 &&
		point.y <= rectangle.y2, nil
}

func (t *TSolver) HouseSurfaceArea(height float64, width float64, windowHeight float64, windowWidth float64) float64 {
	if height <= 0 || width <= 0 || windowWidth <= 0 || windowHeight <= 0 {
		return -1
	}
	var windowsArea = 4 * 2 * windowHeight * windowWidth
	var surface = 5*height*width - windowsArea
	return surface
}

func (t *TSolver) ThreeNumsSum(num1 float64, num2 float64, num3 float64) float64 {
	return num1 + num2 + num3
}

func (t *TSolver) ThreeNumsMin(num1 float64, num2 float64, num3 float64) float64 {
	if num1 <= num2 && num1 <= num3 {
		return num1
	} else if num2 <= num1 && num2 <= num3 {
		return num2
	} else {
		return num3
	}
}
