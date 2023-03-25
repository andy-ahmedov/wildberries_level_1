package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func Distance(p1 *Point, p2 *Point) float64 {
	result := math.Sqrt(math.Pow((p2.x-p1.x), 2) + math.Pow((p2.y-p1.y), 2))
	return result
}

func main() {
	p1 := NewPoint(2, 3)
	p2 := NewPoint(5, 7)

	dist := Distance(p1, p2)
	fmt.Println(dist)

}
