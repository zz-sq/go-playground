package main

import (
	"fmt"
	"image/color"
)

type Point struct{ X, Y float64 }

func (p Point) Print() {
	fmt.Printf("x: %.2f y: %.3f\n", p.X, p.Y)
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	cp.Y = 2
	cp.Color = color.RGBA{255, 0, 0, 255}
	cp.Print()
	fmt.Println(cp.Point)
}
