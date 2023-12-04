package main

import (
	"fmt"
)

type engine struct {
	electric bool
}

type vehicle struct {
	engine engine
	make   string
	model  string
	doors  int
	color  string
}

func main() {

	v1 := vehicle{
		engine: engine{electric: false},
		make:   "Seat",
		model:  "Leon",
		doors:  4,
		color:  "red",
	}
	v2 := vehicle{
		engine: engine{electric: true},
		make:   "Tesla",
		model:  "Y",
		doors:  4,
		color:  "blue",
	}

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println("------------")
	fmt.Println("Colour v1: " + v1.color)
	fmt.Println("Colour v2: " + v2.color)

}

/*
Create a type engine struct, and include this field
○ electric bool
● Create a type vehicle struct, and include these fields
■ engine
■ make
■ model
■ doors
■ color
● Create two VALUES of TYPE vehicle
○ use a composite literal
● Print out each of these values.
● Print out a single field from each of these values.
*/
