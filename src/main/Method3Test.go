package main

import (
	"math"
	"fmt"
)

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(math.Pow(p.x, 2) + math.Pow(p.y, 2))
}

type NamedPoint struct {
	Point
	name string
}

func (n *NamedPoint) Abs() float64 {
	return n.Point.Abs() * 100
}

func main() {
	n := &NamedPoint{Point{3, 4}, "Pythagoras"}
	fmt.Println(n.Point.Abs())
	fmt.Println(n.Abs())
}
