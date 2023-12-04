package main

import "fmt"

func main() {
	var name string = "Peter"
	age := 60
	var pi float64 = 3.14

	fmt.Printf("\n%v\t%T\n", name, name)
	fmt.Printf("\n%v\t%T\n", age, age)
	fmt.Printf("\n%v\t%T\n", pi, pi)

}

/*
write a program that uses the following:
● for a VARIABLE storing a VALUE of TYPE
	○ string
	○ int
	○ float64
● use print verbs to show
	○ value
	○ type
*/
