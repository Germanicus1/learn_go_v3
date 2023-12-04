package main

import (
	"fmt"
)

func main() {

	type person struct {
		first    string
		last     string
		iceCream []string
	}

	p1 := person{
		first:    "James",
		last:     "Bond",
		iceCream: []string{"Chocolate", "strawberry", "vanilla"},
	}
	p2 := person{
		first:    "Jenny",
		last:     "Moneypenny",
		iceCream: []string{"coco", "melon", "peach"},
	}

	fmt.Println(p1.first, p1.last)
	for _, k := range p1.iceCream {
		fmt.Println(k)
	}
	fmt.Println(p2.first, p2.last)
	for _, k := range p2.iceCream {
		fmt.Println(k)
	}

}

/*
Create your own type “person” which will have an underlying type of “struct” so that it can store
the following data:
● first name
● last name
● favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
which stores the favorite flavors.

*/
