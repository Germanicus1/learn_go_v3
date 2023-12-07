package main

import (
	"fmt"
	"math"
)

func main() {
	outer := powerater(3)

	fmt.Println(outer())
	fmt.Println(outer())
	fmt.Println(outer())
	fmt.Println(outer())
	fmt.Println(outer())
}

func powerater(n float64) func() float64 {
	var x float64
	return func() float64 {
		x++
		return math.Pow(n, x)
	}
}

/*
Closure is when we have “enclosed” the scope of a variable in some code block. For this
hands-on exercise, create a func which “encloses” the scope of a variable:

*/
