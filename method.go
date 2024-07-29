package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model string
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := Circle{x: 10, y: 10, r: 10}
	fmt.Println(c.area())

	p := Person{"Saurabh"}
	p.Talk()

	a := new(Android)
	a.Talk()
}
