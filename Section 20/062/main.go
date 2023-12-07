package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func info(s Shape) float64 {
	return s.Area()
}

func main() {
	r := Rectangle{length: 10, width: 20}
	c := Circle{radius: 10}

	fmt.Println("Rectangle area:", info(r))
	fmt.Println("Circle area:", info(c))
}

/*
● create a type SQUARE
	○ length float64
	○ width float64
● create a type CIRCLE
	○ radius float64
● attach a method to each that calculates AREA and returns it
	○ circle area= π r 2
		■ math.Pi
		■ math.Pow
	○ square area = L * W
● create a type SHAPE that defines an interface as anything that has the AREA method
● create a func INFO which takes type shape and then prints the area
● create a value of type square
● create a value of type circle
● use func info to print the area of square
● use func info to print the area of circle
*/
