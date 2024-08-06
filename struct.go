package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := Circle{x: 10, y: 55, r: 10}
	fmt.Println(circleArea(c))
}
