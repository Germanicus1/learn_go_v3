package main

import "fmt"

type person struct {
	first string
}

func changeFirst(p person, s string) person {
	p.first = s
	return p
}

func changeFirstPtr(p *person, s string) {
	p.first = s
}

func main() {
	p1 := person{first: "Peter"}
	p1 = changeFirst(p1, "Hugo")
	fmt.Println(p1.first)
	changeFirstPtr(&p1, "Carmen")
	fmt.Println(p1.first)
}

/*
Create two functions to change a field in a struct called "first" of TYPE string:
● One function will use VALUE SEMANTICS
	○ this function will return some VALUE of some TYPE
● The other function will use POINTER SEMANTICS
	○ this function will return nothing
● don't use methods
*/
