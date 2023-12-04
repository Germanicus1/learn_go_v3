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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	// fmt.Println(p1.first, p1.last)
	// for _, k := range p1.iceCream {
	// 	fmt.Println(k)
	// }
	// fmt.Println(p2.first, p2.last)
	// for _, k := range p2.iceCream {
	// 	fmt.Println(k)
	// }

}

/*
Take the code from the previous exercise, then store the VALUES of type person in a map with the KEY of last name. Access each value in the map. Print out the values, ranging over the slice

*/
